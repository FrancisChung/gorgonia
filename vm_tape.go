package gorgonia

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/chewxy/hm"
	"github.com/pkg/errors"
	"github.com/xtgo/set"
	"gorgonia.org/tensor"
)

type tapeMachine struct {
	ExternMetadata

	p      *program
	locMap map[*Node]register

	// "register" banks
	cpumem   []Value // Value - knows its own type and shape
	gpumem   []Value // Value of which the memories are stored in GPU memory
	cpulocks []sync.RWMutex
	gpulocks []sync.RWMutex

	// state stuff, to allow continuation after failure handling
	pc int

	// operational stuff
	bindNodesDV Nodes // nodes that require binding of DV
	watchNodes  Nodes
	watchRegs   []register
	logger      *log.Logger
	buf         *bytes.Buffer
	valueFmt    string
	tabcount    int
	logFlags    byte

	sync.Mutex

	runFlags byte //  spare2: trace (copy values and put into nodes)
}

// NewTapeMachine creates a VM that compiles a graph into a prog.
func NewTapeMachine(g *ExprGraph, opts ...VMOpt) *tapeMachine {
	m := &tapeMachine{
		valueFmt: "%3.3g",
	}
	m.Engine = StandardEngine{}

	if b, ok := whichblas.(batchedBLAS); ok {
		m.b = b
	}

	for _, opt := range opts {
		opt(m)
	}

	m.doAlloc()

	if m.p == nil || m.locMap == nil {
		prog, locMap, err := Compile(g)
		if err != nil {
			panic(err)
		}

		m.p = prog
		m.locMap = locMap
		m.cpumem = make([]Value, prog.cpulocs)
		m.gpumem = make([]Value, prog.gpulocs)
		m.cpulocks = make([]sync.RWMutex, prog.cpulocs)
		m.gpulocks = make([]sync.RWMutex, prog.gpulocs)
	}
	m.init() // init ExternalMetadata
	for _, n := range m.p.g.AllNodes() {
		setEngine(n.boundTo, m.Engine)
	}

	runtime.SetFinalizer(m, finalizeTapeMachine) // a "defer" to deinitialize CUDA stuff (if using CUDA build)
	return m
}

func (m *tapeMachine) logBwd() bool { return (m.logFlags>>bwdOnly)&byte(1) == 1 }
func (m *tapeMachine) doLogBwd()    { m.logFlags |= byte(1) << bwdOnly }
func (m *tapeMachine) dontLogBwd()  { m.logFlags &= (^(byte(1) << bwdOnly)) }

func (m *tapeMachine) logFwd() bool { return (m.logFlags>>fwdOnly)&byte(1) == 1 }
func (m *tapeMachine) doLogFwd()    { m.logFlags |= byte(1) << fwdOnly }
func (m *tapeMachine) dontLogFwd()  { m.logFlags &= (^(byte(1) << fwdOnly)) }

func (m *tapeMachine) watchNaN() bool { return (m.runFlags>>watchNaN)&byte(1) == 1 }
func (m *tapeMachine) doWatchNaN()    { m.runFlags |= byte(1) << watchNaN }
func (m *tapeMachine) dontWatchNaN()  { m.runFlags &= (^(byte(1) << watchNaN)) }

func (m *tapeMachine) watchInf() bool { return (m.runFlags>>watchInf)&byte(1) == 1 }
func (m *tapeMachine) doWatchInf()    { m.runFlags |= byte(1) << watchInf }
func (m *tapeMachine) dontWatchInf()  { m.runFlags &= (^(byte(1) << watchInf)) }

func (m *tapeMachine) watchAll() bool { return (m.logFlags>>watchAll)&byte(1) == 1 }
func (m *tapeMachine) doWatchAll()    { m.logFlags |= (byte(1) << watchAll) }
func (m *tapeMachine) dontWatchAll()  { m.logFlags &= (^(byte(1) << watchAll)) }

func (m *tapeMachine) alloc() bool { return (m.runFlags>>allocVals)&byte(1) == 1 }
func (m *tapeMachine) doAlloc()    { m.runFlags |= byte(1) << allocVals }
func (m *tapeMachine) dontAlloc()  { m.runFlags &= (^(byte(1) << allocVals)) }

func (m *tapeMachine) trace() bool { return (m.runFlags>>spare2)&byte(1) == 1 }
func (m *tapeMachine) doTrace()    { m.runFlags |= byte(1) << spare2 }
func (m *tapeMachine) dontTrace()  { m.runFlags &= (^(byte(1) << spare2)) }

func (m *tapeMachine) bindDV() bool { return m.runFlags>>spare3&byte(1) == 1 }
func (m *tapeMachine) doBindDV()    { m.runFlags |= byte(1) << spare3 }
func (m *tapeMachine) dontBindDV()  { m.runFlags &= (^(byte(1) << spare3)) }

// Reset resets the run state of the machine by changing the instruction pointer back to 0
func (m *tapeMachine) Reset() {
	m.pc = 0
	m.ExternMetadata.Reset()
}

// Prog returns the compiled program. This would mainly be used in debugging functions
func (m *tapeMachine) Prog() *program { return m.p }

// LocMap returns the location where the Node's execution results are stored. This would mainly be used in debugging functions.
func (m *tapeMachine) LocMap() map[*Node]register { return m.locMap }

// Let wraps the Let() function of the package, with additional checks that n is in the machine
func (m *tapeMachine) Let(n *Node, be interface{}) (err error) {
	if !m.p.g.Has(n.ID()) {
		return errors.Errorf("Node %v does not exist in this graph", n)
	}

	return Let(n, be)
}

// Set wraps the Set() function of this package, with additional checks that both a and b are in the machine
func (m *tapeMachine) Set(a, b *Node) (err error) {
	if !m.p.g.Has(a.ID()) {
		return errors.Errorf("Node %v does not exist in this graph", a)
	}
	if !m.p.g.Has(b.ID()) {
		return errors.Errorf("Node %v does not exist in this graph", b)
	}

	if b.Value() != nil {
		return a.bind(b.Value())
	}

	// get the registry location
	breg := m.locMap[b]
	v := m.getValue(breg)
	if v == nil {
		return nyi("handling of tensor.Memory -> Value", "tapeMachine.Set")
	}

	machineLogf("Setting %v to %v. Read from %v Value is %v", b, a, breg, v)
	return a.bind(v)
}

// Run runs a fragment (a subset of a program).
func (m *tapeMachine) Run(frag fragment) (err error) {
	defer func() {
		if err == nil {
			m.dontAlloc()
		}
	}()

	for _, instr := range frag {
		if err = instr.exec(m); err != nil {
			return errors.Wrap(err, "Failed to carry exec()")
		}
	}
	machineLogf("Binding values based on final output")
	enterLogScope()
	for n, r := range m.locMap {
		if n.isInput() {
			continue
		}

		v := m.getValue(r)
		if v == nil {
			return nyi("converting tensor.Memory to Value", "TapeMachine.Run")
		}

		if err = n.bind(m.cpumem[r.id]); err != nil {
			return errors.Wrap(err, bindFail)
		}
	}
	leaveLogScope()
	return
}

func (m *tapeMachine) RunAll() (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	workAvailable := m.ExternMetadata.WorkAvailable()
	syncChan := m.ExternMetadata.Sync()
	errChan := make(chan error)
	doneChan := make(chan struct{})

	go m.runall(errChan, doneChan)
	for {
		select {
		case sychronous := <-workAvailable:
			err := m.ExternMetadata.DoWork()
			if err != nil {
				return err
			}
			if sychronous {
				syncChan <- struct{}{}
			}
		case err := <-errChan:
			return errors.Wrapf(err, "PC: %d", m.pc)
		case <-doneChan:
			err := m.ExternMetadata.DoWork()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return
}

var print11 bool

func (m *tapeMachine) runall(errChan chan error, doneChan chan struct{}) {
	m.buf = new(bytes.Buffer)
	s := newExecState(m.p.g, m.p.sorted, m.p, m.buf)
	var wg sync.WaitGroup
	threads := runtime.NumCPU()
	workers := make(chan struct{}, threads)
	for t := 0; t < threads; t++ {
		wg.Add(1)
		go m.execute(s, workers, errChan, &wg, t)
	}
	wg.Wait()
	// log.Println(m.buf.String())
	doneChan <- struct{}{}
}

func (m *tapeMachine) execute(s *execState, workers chan struct{}, errChan chan error, wg *sync.WaitGroup, id int) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

execloop:
	for {
		// <-workers
		if s.check() {
			break execloop
		}

		pnode := s.next()
		if pnode == nil {
			// workers <- struct{}{}
			continue
		}
		instrs := m.p.m[pnode.Node]
		for _, instr := range instrs {

			fmt.Fprintf(m.buf, "Executing %d : %v | %v\n", pnode.index, pnode, instr)
			// log.Printf("Executing %d %v\n", pnode.index, pnode)

			if err := m.executeOneInstr(instr); err != nil {
				err = errors.Wrapf(err, "pnode %d: %v", pnode.index, pnode)
				errChan <- err
				s.error()
				// workers <- struct{}{}
				break execloop
			}
		}
		s.finish(pnode)
		// workers <- struct{}{}
	}
	wg.Done()
}

func (m *tapeMachine) executeOneInstr(instr tapeInstr) error {
	if err := instr.exec(m); err != nil {
		return errors.Wrapf(err, "Failed to execute instruction %v", instr)
	}

	if m.watchNaN() {
		writeTo := instr.writes().id
		id := instr.ID()
		if writeTo > 0 && id > 0 {
			v := m.getValue(instr.writes())
			if v == nil {
				return errors.Errorf(nyiFail, "converting tensor.Memory to Value", "watchNaN")
			}

			if hasNaN(v) {
				n := m.p.g.Node(id).(*Node)
				return errors.Errorf("NaN found in value. Node: %v(%x)", n, n.ID())
			}
		}
	}

	if m.watchInf() {
		writeTo := instr.writes().id
		id := instr.ID()
		if writeTo > 0 && id > 0 {
			v := m.getValue(instr.writes())
			if v == nil {
				return errors.Errorf(nyiFail, "converting tensor.Memory to Value", "watchInf")
			}

			if hasInf(v) {
				n := m.p.g.Node(id).(*Node)
				return errors.Errorf("Inf found in value. Node: %v(%x)", n, n.ID())
			}
		}
	}
	return nil
}

func (m *tapeMachine) getValue(r register) Value {
	switch r.device {
	case CPU:
		return m.cpumem[r.id]
	default:
		return m.gpumem[r.id]
	}
}

func (m *tapeMachine) writeValue(r register, v Value) {
	switch r.device {
	case CPU:
		m.cpumem[r.id] = v
	default:
		m.gpumem[r.id] = v
	}
}

func (m *tapeMachine) watchedInstrLogf(instr tapeInstr, format string, attrs ...interface{}) {
	reads := instr.reads()
	writes := instr.writes()
	watched := m.watchAll()

	if !watched {
		for _, reg := range reads {
			for _, watch := range m.watchRegs {
				if reg.id == watch.id {
					watched = true
					break
				}
			}
		}
	}

	if !watched {
		for _, watch := range m.watchRegs {
			if watch.id == writes.id {
				watched = true
				break
			}
		}
	}

	// TODO: Work on watched nodes
	if !watched {

	}

	if watched {
		m.logf(format, attrs...)
	}
}

func (m *tapeMachine) logf(format string, attrs ...interface{}) {
	switch {
	case machineDev:
		if m.logger != nil {
			goto loggercase
		}

		machineLogf(format, attrs...)
		break

	loggercase:
		fallthrough
	case m.logger != nil:
		s := fmt.Sprintf(format, attrs...)
		s = strings.Replace(s, "\n", m.buf.String(), -1)
		m.logger.Println(s)
	}
}

func (m *tapeMachine) enterLogScope() {
	if DEBUG && machineDev {
		enterLogScope()
	}
	m.tabcount++
	if m.logger != nil {
		reps := strings.Repeat("\t", m.tabcount)
		m.logger.SetPrefix(reps)
		m.buf.Reset()
		m.buf.WriteString("\n")
		m.buf.WriteString(reps)
	}
}

func (m *tapeMachine) leaveLogScope() {
	if DEBUG && machineDev {
		leaveLogScope()
	}
	m.tabcount--
	if m.tabcount < 0 {
		m.tabcount = 0
	}
	if m.logger != nil {
		reps := strings.Repeat("\t", m.tabcount)
		m.logger.SetPrefix(reps)
		m.buf.Reset()
		m.buf.WriteString("\n")
		m.buf.WriteString(reps)
	}
}

/* PROGRAM */

type program struct {
	instructions fragment
	args         int
	cpulocs      int
	gpulocs      int
	cpumem       int64
	gpumem       []int64
	g            *ExprGraph         // original dag
	df           *dataflow          // dataflow analysis
	m            map[*Node]fragment // store which nodes create which instructions
	r            map[*Node]intervalRange
	sorted       Nodes
}

func (p *program) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Instructions:\n%s\nArgs: %d | CPU Memories: %d | GPU Memories: %d\nCPU Mem: %v | GPU Mem %v\n\nNode:instructions map:\n", p.instructions, p.args, p.cpulocs, p.gpulocs, p.cpumem, p.gpumem)

	for i, n := range p.sorted {
		fmt.Fprintf(&buf, "\t%d\t%x:", i, n.ID())
		frag := p.m[n]
		for j, instr := range frag {
			if j == 0 {
				fmt.Fprintf(&buf, "\t%v\n", instr)
			} else {
				fmt.Fprintf(&buf, "\t\t%v\n", instr)
			}
		}

	}

	return buf.String()
}

// Graph enables the end user to inspect the graph (typically useful for debugging)
func (p *program) Graph() *ExprGraph { return p.g }

func (p *program) CPUMemReq() int64 { return p.cpumem }

func (p *program) GPUMemReq() []int64 {
	retVal := make([]int64, len(p.gpumem))
	copy(retVal, p.gpumem)
	return retVal
}

/* REGISTER */

type register struct {
	id     int
	device Device
}

func (r register) String() string { return fmt.Sprintf("%s%d", r.device, r.id) }

/* INSTRUCTIONS */

type tapeInstr interface {
	ID() int64 // ID is the node ID
	reads() []register
	writes() register
	exec(*tapeMachine) error
	fmt.Stringer
}

type fragment []tapeInstr

func (f fragment) String() string {
	var buf bytes.Buffer
	for i, instr := range f {
		fmt.Fprintf(&buf, "\t%d\t%s\n", i, instr)
	}
	return buf.String()
}

func (f fragment) has(want tapeInstr) bool {
	for _, instr := range f {
		if instr == want {
			return true
		}
	}
	return false
}

type alloc struct {
	id int64 // node ID
	t  hm.Type
	s  tensor.Shape

	readFrom []register
	writeTo  register
}

func newAlloc(n *Node, writeTo register) alloc {
	return alloc{
		id:      n.ID(),
		t:       n.t,
		s:       n.shape,
		writeTo: writeTo,
	}
}

func (instr alloc) ID() int64         { return instr.id }
func (instr alloc) reads() []register { return instr.readFrom }
func (instr alloc) writes() register  { return instr.writeTo }

func (instr alloc) exec(m *tapeMachine) (err error) {
	m.logf("Executing %v", instr)

	var dt tensor.Dtype
	if dt, err = dtypeOf(instr.t); err != nil {
		return errors.Wrapf(err, dtypeExtractionFail, instr.t)
	}

	dev := instr.writeTo.device
	var v Value
	switch dev {
	case CPU:
		v, err = makeValue(instr.t, instr.s)
	default:
		var mem tensor.Memory
		memsize := calcMemSize(dt, instr.s)
		if mem, err = m.ExternMetadata.Get(dev, memsize); err != nil {
			return errors.Wrapf(err, "Unable to allocate %v bytes from %v", memsize, dev)
		}
		v, err = makeValueFromMem(instr.t, instr.s, mem)
	}
	if err != nil {
		return
	}

	m.writeValue(instr.writeTo, v)
	return nil
}

func (instr alloc) String() string {
	return fmt.Sprintf("Alloc %v%v\t\t%v", instr.t, instr.s, instr.writeTo)
}

type free struct {
	readsFrom register
}

func (instr free) ID() int64         { return -1 }
func (instr free) reads() []register { return []register{instr.readsFrom} }
func (instr free) writes() register  { return register{-1, CPU} }
func (instr free) exec(m *tapeMachine) error {
	m.logf("Executing Free %v", instr.readsFrom)
	switch instr.readsFrom.device {
	case CPU:
		return nil
	default:
		m.logf("instr.read from not CPU - %v %v %d", instr.readsFrom, instr.readsFrom.device == CPU, instr.readsFrom.device)
		mem := m.gpumem[instr.readsFrom.id]
		size := int64(mem.MemSize())

		m.Put(instr.readsFrom.device, mem, size)
		m.gpumem[instr.readsFrom.id] = nil
		return nil
	}
}
func (instr free) String() string { return fmt.Sprintf("Free %v", instr.readsFrom) }

type loadArg struct {
	index   int64
	writeTo register
}

func (instr loadArg) ID() int64         { return instr.index }
func (instr loadArg) reads() []register { return nil }
func (instr loadArg) writes() register  { return instr.writeTo }

func (instr loadArg) exec(m *tapeMachine) error {
	m.logf("Executing %v", instr)
	m.enterLogScope()
	defer m.leaveLogScope()

	node := m.p.g.Node(instr.index).(*Node)
	m.logf("node %v", node)

	if node.boundTo == nil {
		return errors.Errorf("No value bound to node %v (%x)", node, node.ID())
	}

	var v Value
	if dv, ok := node.boundTo.(*dualValue); ok {
		v = dv.Value
	} else {
		v = node.boundTo
	}

	m.writeValue(instr.writeTo, v)
	return nil
}

func (instr loadArg) String() string {
	return fmt.Sprintf("loadArg %x to %v", instr.index, instr.writeTo)
}

type execOp struct {
	op Op

	id int64

	readFrom []register
	writeTo  register
	size     int64 // size represents the outputsize

	preAllocated bool
	useUnsafe    bool
	useGPU       bool
}

func (instr *execOp) ID() int64         { return instr.id }
func (instr *execOp) reads() []register { return instr.readFrom }
func (instr *execOp) writes() register  { return instr.writeTo }

func newExecOp(n *Node) *execOp {
	_, useGPU := n.op.(CUDADoer)
	compileLogf("op %v uses GPU %v", n.op, useGPU)
	dt, err := dtypeOf(n.t)
	if err != nil {
		panic(err)
	}
	size := calcMemSize(dt, n.Shape())

	return &execOp{
		op:     n.op,
		id:     n.ID(),
		useGPU: useGPU,
		size:   size,
	}
}

func (instr *execOp) exec(m *tapeMachine) error {
	m.logf("Executing %v. Node is: %x", instr, instr.id)
	m.enterLogScope()
	defer m.leaveLogScope()

	// Read
	m.watchedInstrLogf(instr, "Inputs:")
	m.enterLogScope()
	inputs := make([]Value, 0, len(instr.readFrom))
	for _, reg := range instr.readFrom {
		v := m.getValue(reg)
		inputs = append(inputs, v)
		m.watchedInstrLogf(instr, m.valueFmt, v)
	}
	m.leaveLogScope()

	err := instr.execKernel(m, inputs)
	return err
}

func (instr *execOp) String() string {
	return fmt.Sprintf("%v\t%v\t%v\t%t\t%t\t%t", instr.op, instr.readFrom, instr.writeTo, instr.op.CallsExtern(), instr.useUnsafe, instr.preAllocated)
}

// flushInstr is for blastoise and cubone
type flushInstr struct{}

func (instr flushInstr) exec(m *tapeMachine) error {
	// m.logf("Executing DoWork")
	// m.ExternMetadata.DoWork()
	if m.WorkAvailable() == nil {
		return nil
	}
	m.ExternMetadata.Signal()
	return nil
}

func (instr flushInstr) ID() int64         { return -1 }
func (instr flushInstr) reads() []register { return nil }
func (instr flushInstr) writes() register  { return register{-1, CPU} }
func (instr flushInstr) String() string    { return "DoWork" }

type letInstr struct {
	readFrom register
	writeTo  register
}

func (instr letInstr) ID() int64               { return -1 }
func (instr letInstr) reads() []register       { return []register{instr.readFrom} }
func (instr letInstr) writes() register        { return instr.writeTo }
func (instr letInstr) exec(*tapeMachine) error { return nil }

func (instr letInstr) String() string {
	return fmt.Sprintf("LET %v = %v", instr.writeTo, instr.readFrom)
}

type readInstr struct {
	readFrom register
	into     *Value

	// required to convert tensor.Memory to Value
	t hm.Type
	s tensor.Shape
}

func (instr *readInstr) ID() int64         { return -1 }
func (instr *readInstr) reads() []register { return []register{instr.readFrom} }
func (instr *readInstr) writes() register  { return register{-1, CPU} }
func (instr *readInstr) exec(m *tapeMachine) (err error) {
	m.logf("Executing READ - read from %v into %v", instr.readFrom, instr.into)
	v := m.getValue(instr.readFrom)
	if v == nil {
		return nyi("value of nil", "readInstr.exec")
	}

	v2, err := CloneValue(v)
	if err != nil {
		return errors.Wrap(err, cloneFail)
	}

	*instr.into = v2
	m.logf("instr.into %v", *instr.into)
	return nil
}

func (instr *readInstr) String() string {
	return fmt.Sprintf("Read %v into %p", instr.readFrom, instr.into)
}

type deviceTransport struct {
	from, to register
}

func (instr deviceTransport) ID() int64 { return -1 }
func (instr deviceTransport) reads() []register {
	return []register{instr.from}
}
func (instr deviceTransport) writes() register { return instr.to }

func (instr deviceTransport) String() string {
	return fmt.Sprintf("memcpy(%v, %v)", instr.to, instr.from)
}

type execState struct {
	sync.RWMutex
	sorted    []priorityNode
	p         *program
	m         map[*Node]int
	q         []int // queue where all the dependents are ready
	t         [][]int
	f         [][]int
	r         map[register][]int
	w         map[register][]int
	cpuWrites []int32
	gpuWrites []int32

	donecount int
	workers   int
	nodes     int
	done      bool
	err       bool

	buf *bytes.Buffer
}

func newExecState(g *ExprGraph, sorted Nodes, p *program, buf *bytes.Buffer) *execState {
	s := make([]priorityNode, len(sorted))
	m := make(map[*Node]int, len(sorted))
	for i := range s {
		n := sorted[i]
		s[i] = priorityNode{
			Node:     n,
			priority: int32(len(n.children)),
			index:    i,
		}
		m[n] = i
	}

	// "lift" tos and froms
	t := make([][]int, len(s))
	f := make([][]int, len(s))
	for k, v := range g.to {
		id := m[k]
		t[id] = make([]int, 0, len(v)+4) // 4 is spare
		for _, n := range v {
			nid := m[n]
			t[id] = append(t[id], nid)
		}
	}
	for _, n := range s {
		id := n.index
		f[id] = make([]int, 0, len(n.children)+4) // 4 is spare
		for _, child := range n.children {
			nid := m[child]
			f[id] = append(f[id], nid)
		}
	}

	reads := make(map[register][]int)
	writes := make(map[register][]int)

	for i := range s {
		pn := &s[i]
		n := pn.Node
		instrs := p.m[n]
		for _, instr := range instrs {
			for _, r := range instr.reads() {
				reads[r] = append(reads[r], i)
				// pn.reads.Add(r)
			}
			w := instr.writes()
			writes[w] = append(writes[w], i)
			// pn.writes.Add(w)

			if ri, ok := instr.(*readInstr); ok {
				writes[ri.readFrom] = append(writes[ri.readFrom], i)
			}
		}
	}

	// uniquify the reads and writes
	for k, v := range reads {
		reads[k] = set.Ints(v)
	}
	for k, v := range writes {
		writes[k] = set.Ints(v)
	}
	/*
		for i := range s {
			pn := &s[i]
			n := pn.Node
			instrs := p.m[n]
			for _, instr := range instrs {
				if _, ok := instr.(free); ok {
					continue
				}

				for _, r := range instr.reads() {
					rs := reads[r]
					var atI int
					for j := len(rs) - 1; j >= 0; j-- {
						nid := rs[j]
						if nid > i {
							continue
						}
						if nid == i {
							atI = j
							continue
						}

						log.Printf("\tnid %v, i %d", nid, i)

						for _, wid := range writes[r] {
							var hasReadInstr bool
							log.Printf("\t\t%v", s[wid].Node)
							for _, in := range p.m[s[wid].Node] {
								// log.Printf("\t\t\t%v", in)
								if _, ok := in.(*readInstr); ok {
									log.Printf("\t\tHAS READ")
									hasReadInstr = true
									break
								}
							}

							// if the writing instruction is the same as the reading instruction, we need to actually add a dependency
							if wid == i && (hasReadInstr || j == atI-1) {
							// if wid == i {
								log.Printf("\t\tadding %d | %d | %v %v", nid, i, hasReadInstr, atI)
								pn.priority++
								t[nid] = append(t[nid], i)
								f[i] = append(f[i], nid)
							}
						}
					}
					// var j, nid int
					// for j, nid = range reads[r] {
					// 	if nid == i {
					// 		break
					// 	}
					// }
					// if j == 0 {
					// 	break
					// }
					// nid = reads[r][j-1]
					// if nid > i {
					// break
					// }

					// log.Printf("r %v nid %d, adding %d | %v | %v", r, nid, i, reads[r], writes[r])
					// check if the instruction is writing the register

					// t[nid] = append(t[nid], i)
				}
			}
		}
	*/

	for reg, wnids := range writes {
		if reg.id == -1 {
			continue
		}
		// log.Printf("nodes that write %v | %v", reg, wnids)
		for _, nid := range wnids {
			rnids := reads[reg]
			for _, rid := range rnids {
				if rid >= nid {
					continue
				}
				// log.Printf("\t%d reads %v: %v | %v", nid, reg, rid, t[rid])
				s[nid].priority++
				f[nid] = append(f[nid], rid)
				t[rid] = append(t[rid], nid)
			}
		}
	}

	for i, v := range t {
		t[i] = set.Ints(v)
	}

	// var buf bytes.Buffer
	fmt.Fprintf(buf, "PROG\n%v", p)
	fmt.Fprintf(buf, "SORTED:\n")
	for i, v := range s {
		fmt.Fprintf(buf, "%d: %v\n", i, v)
	}
	fmt.Fprintf(buf, "To\n")
	for i, v := range t {
		fmt.Fprintf(buf, "%d: %v\n", i, v)
	}
	fmt.Fprintf(buf, "From\n")
	for i, v := range f {
		fmt.Fprintf(buf, "%d: %v\n", i, v)
	}
	// log.Printf("%v", buf.String())

	// prefill the queue
	q := make([]int, 0, len(g.leaves)+len(g.constants))
	for _, leaf := range g.leaves {
		q = append(q, m[leaf])
	}
	for _, c := range g.constants {
		q = append(q, m[c])
	}

	return &execState{
		sorted: s,
		p:      p,
		m:      m,
		q:      q,
		t:      t,
		f:      f,
		r:      reads,
		w:      writes,
		nodes:  len(sorted),

		buf: buf,
	}
}

func (s *execState) finish(node *priorityNode) {
	// log.Printf("FINISHED %x: %v", node.id, node)
	s.Lock()
	fmt.Fprintf(s.buf, "FINISHED %d: %v\n", node.index, node)
	to := s.t[node.index]
	s.Unlock()

	for _, i := range to {
		n := &s.sorted[i]
		// this loop is necessary because to only records a single edge. There may be multiple edges to the same node
		var reduction int32
		for _, j := range s.f[i] {
			if j == node.index {
				reduction--
			}
		}

		toNodePriority := atomic.AddInt32(&n.priority, reduction)
		if toNodePriority == 0 {
			s.Lock()
			if !n.finished { // this line shouldn't be necessary
				// log.Printf("\t %v added to queue", n)
				s.q = append(s.q, n.index)
			}
			s.Unlock()
		}
		// } else {
		// 	log.Printf("\t%v | %d", n, toNodePriority)
		// }
	}
	// var nodes, workers, q int
	s.Lock()
	node.finished = true
	s.donecount++
	s.nodes--
	s.workers--
	if s.nodes == 0 && s.workers == 0 && s.donecount == len(s.sorted) {
		s.done = true
	}
	s.Unlock()
}

func (s *execState) next() (retVal *priorityNode) {
	s.Lock()
	if len(s.q) == 0 {
		goto end
	}
	if len(s.q) > 1 {
		sort.Slice(s.q, func(i, j int) bool { return s.sorted[s.q[i]].index > s.sorted[s.q[j]].index })
	}
	fmt.Fprintf(s.buf, "Next. State: %v\n", s.q)
	retVal = &s.sorted[s.q[len(s.q)-1]]
	s.q = s.q[:len(s.q)-1]
	// for i := 0; i < len(s.q); i++ {
	// 	n := s.q[i]
	// 	if s.checkRegisterDependency(&s.sorted[n]) {
	// 		s.q2 = append(s.q2, n)
	// 		copy(s.q[i:], s.q[i+1:])
	// 		s.q[len(s.q)-1] = 0
	// 		s.q = s.q[:len(s.q)-1]
	// 		i--
	// 	}
	// }
	// if len(s.q2) == 0 {
	// 	goto end
	// }

	// if len(s.q2) >= 2 {
	// 	sort.Slice(s.q2, func(i, j int) bool { return s.sorted[s.q2[i]].index > s.sorted[s.q2[j]].index })
	// }
	// retVal = &s.sorted[s.q2[len(s.q2)-1]]
	// s.q2 = s.q2[:len(s.q2)-1]
	s.workers++
end:
	s.Unlock()
	return retVal
}

func (s *execState) error() {
	s.Lock()
	s.err = true
	s.Unlock()
}

func (s *execState) check() bool {
	s.RLock()
	retVal := s.done || s.err
	s.RUnlock()
	return retVal
}

type priorityNode struct {
	*Node
	priority  int32
	priority2 int32
	index     int
	finished  bool

	reads  registerSet
	writes registerSet
}

type registerSet map[register]struct{}

func makeRegisterSet() registerSet { return make(map[register]struct{}) }

func (s registerSet) Add(r register) registerSet {
	s[r] = struct{}{}
	return s
}

func (s registerSet) Contains(r register) bool {
	_, ok := s[r]
	return ok
}
