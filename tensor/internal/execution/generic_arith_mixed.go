package execution

import (
	"math"
	"math/cmplx"

	"github.com/chewxy/math32"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func AddSVI(a int, b []int) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVI8(a int8, b []int8) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVI16(a int16, b []int16) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVI32(a int32, b []int32) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVI64(a int64, b []int64) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVU(a uint, b []uint) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVU8(a uint8, b []uint8) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVU16(a uint16, b []uint16) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVU32(a uint32, b []uint32) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVU64(a uint64, b []uint64) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVF32(a float32, b []float32) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVF64(a float64, b []float64) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVC64(a complex64, b []complex64) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVC128(a complex128, b []complex128) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func AddSVStr(a string, b []string) {
	for i := range b {
		b[i] = a + b[i]
	}
}

func SubSVI(a int, b []int) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVI8(a int8, b []int8) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVI16(a int16, b []int16) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVI32(a int32, b []int32) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVI64(a int64, b []int64) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVU(a uint, b []uint) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVU8(a uint8, b []uint8) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVU16(a uint16, b []uint16) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVU32(a uint32, b []uint32) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVU64(a uint64, b []uint64) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVF32(a float32, b []float32) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVF64(a float64, b []float64) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVC64(a complex64, b []complex64) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func SubSVC128(a complex128, b []complex128) {
	for i := range b {
		b[i] = a - b[i]
	}
}

func MulSVI(a int, b []int) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVI8(a int8, b []int8) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVI16(a int16, b []int16) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVI32(a int32, b []int32) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVI64(a int64, b []int64) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVU(a uint, b []uint) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVU8(a uint8, b []uint8) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVU16(a uint16, b []uint16) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVU32(a uint32, b []uint32) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVU64(a uint64, b []uint64) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVF32(a float32, b []float32) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVF64(a float64, b []float64) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVC64(a complex64, b []complex64) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func MulSVC128(a complex128, b []complex128) {
	for i := range b {
		b[i] = a * b[i]
	}
}

func DivSVI(a int, b []int) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVI8(a int8, b []int8) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVI16(a int16, b []int16) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVI32(a int32, b []int32) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVI64(a int64, b []int64) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVU(a uint, b []uint) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVU8(a uint8, b []uint8) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVU16(a uint16, b []uint16) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVU32(a uint32, b []uint32) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVU64(a uint64, b []uint64) (err error) {
	var errs errorIndices
	for i := range b {
		if b[i] == 0 {
			errs = append(errs, i)
			b[i] = 0
			continue
		}
		b[i] = a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivSVF32(a float32, b []float32) {
	for i := range b {
		b[i] = a / b[i]
	}
}

func DivSVF64(a float64, b []float64) {
	for i := range b {
		b[i] = a / b[i]
	}
}

func DivSVC64(a complex64, b []complex64) {
	for i := range b {
		b[i] = a / b[i]
	}
}

func DivSVC128(a complex128, b []complex128) {
	for i := range b {
		b[i] = a / b[i]
	}
}

func PowSVF32(a float32, b []float32) {
	for i := range b {
		b[i] = math32.Pow(a, b[i])
	}
}

func PowSVF64(a float64, b []float64) {
	for i := range b {
		b[i] = math.Pow(a, b[i])
	}
}

func PowSVC64(a complex64, b []complex64) {
	for i := range b {
		b[i] = complex64(cmplx.Pow(complex128(a), complex128(b[i])))
	}
}

func PowSVC128(a complex128, b []complex128) {
	for i := range b {
		b[i] = cmplx.Pow(a, b[i])
	}
}

func ModSVI(a int, b []int) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVI8(a int8, b []int8) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVI16(a int16, b []int16) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVI32(a int32, b []int32) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVI64(a int64, b []int64) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVU(a uint, b []uint) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVU8(a uint8, b []uint8) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVU16(a uint16, b []uint16) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVU32(a uint32, b []uint32) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVU64(a uint64, b []uint64) {
	for i := range b {
		b[i] = a % b[i]
	}
}

func ModSVF32(a float32, b []float32) {
	for i := range b {
		b[i] = math32.Mod(a, b[i])
	}
}

func ModSVF64(a float64, b []float64) {
	for i := range b {
		b[i] = math.Mod(a, b[i])
	}
}

func AddIncrSVI(a int, b []int, incr []int) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVI8(a int8, b []int8, incr []int8) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVI16(a int16, b []int16, incr []int16) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVI32(a int32, b []int32, incr []int32) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVI64(a int64, b []int64, incr []int64) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVU(a uint, b []uint, incr []uint) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVU8(a uint8, b []uint8, incr []uint8) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVU16(a uint16, b []uint16, incr []uint16) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVU32(a uint32, b []uint32, incr []uint32) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVU64(a uint64, b []uint64, incr []uint64) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVF32(a float32, b []float32, incr []float32) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVF64(a float64, b []float64, incr []float64) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVC64(a complex64, b []complex64, incr []complex64) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVC128(a complex128, b []complex128, incr []complex128) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func AddIncrSVStr(a string, b []string, incr []string) {
	for i := range incr {
		incr[i] += a + b[i]
	}
}

func SubIncrSVI(a int, b []int, incr []int) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVI8(a int8, b []int8, incr []int8) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVI16(a int16, b []int16, incr []int16) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVI32(a int32, b []int32, incr []int32) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVI64(a int64, b []int64, incr []int64) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVU(a uint, b []uint, incr []uint) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVU8(a uint8, b []uint8, incr []uint8) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVU16(a uint16, b []uint16, incr []uint16) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVU32(a uint32, b []uint32, incr []uint32) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVU64(a uint64, b []uint64, incr []uint64) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVF32(a float32, b []float32, incr []float32) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVF64(a float64, b []float64, incr []float64) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVC64(a complex64, b []complex64, incr []complex64) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func SubIncrSVC128(a complex128, b []complex128, incr []complex128) {
	for i := range incr {
		incr[i] += a - b[i]
	}
}

func MulIncrSVI(a int, b []int, incr []int) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVI8(a int8, b []int8, incr []int8) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVI16(a int16, b []int16, incr []int16) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVI32(a int32, b []int32, incr []int32) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVI64(a int64, b []int64, incr []int64) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVU(a uint, b []uint, incr []uint) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVU8(a uint8, b []uint8, incr []uint8) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVU16(a uint16, b []uint16, incr []uint16) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVU32(a uint32, b []uint32, incr []uint32) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVU64(a uint64, b []uint64, incr []uint64) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVF32(a float32, b []float32, incr []float32) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVF64(a float64, b []float64, incr []float64) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVC64(a complex64, b []complex64, incr []complex64) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func MulIncrSVC128(a complex128, b []complex128, incr []complex128) {
	for i := range incr {
		incr[i] += a * b[i]
	}
}

func DivIncrSVI(a int, b []int, incr []int) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVI8(a int8, b []int8, incr []int8) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVI16(a int16, b []int16, incr []int16) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVI32(a int32, b []int32, incr []int32) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVI64(a int64, b []int64, incr []int64) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVU(a uint, b []uint, incr []uint) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVU8(a uint8, b []uint8, incr []uint8) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVU16(a uint16, b []uint16, incr []uint16) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVU32(a uint32, b []uint32, incr []uint32) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVU64(a uint64, b []uint64, incr []uint64) (err error) {
	var errs errorIndices
	for i := range incr {
		if b[i] == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a / b[i]
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrSVF32(a float32, b []float32, incr []float32) {
	for i := range incr {
		incr[i] += a / b[i]
	}
}

func DivIncrSVF64(a float64, b []float64, incr []float64) {
	for i := range incr {
		incr[i] += a / b[i]
	}
}

func DivIncrSVC64(a complex64, b []complex64, incr []complex64) {
	for i := range incr {
		incr[i] += a / b[i]
	}
}

func DivIncrSVC128(a complex128, b []complex128, incr []complex128) {
	for i := range incr {
		incr[i] += a / b[i]
	}
}

func PowIncrSVF32(a float32, b []float32, incr []float32) {
	for i := range incr {
		incr[i] += math32.Pow(a, b[i])
	}
}

func PowIncrSVF64(a float64, b []float64, incr []float64) {
	for i := range incr {
		incr[i] += math.Pow(a, b[i])
	}
}

func PowIncrSVC64(a complex64, b []complex64, incr []complex64) {
	for i := range incr {
		incr[i] += complex64(cmplx.Pow(complex128(a), complex128(b[i])))
	}
}

func PowIncrSVC128(a complex128, b []complex128, incr []complex128) {
	for i := range incr {
		incr[i] += cmplx.Pow(a, b[i])
	}
}

func ModIncrSVI(a int, b []int, incr []int) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVI8(a int8, b []int8, incr []int8) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVI16(a int16, b []int16, incr []int16) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVI32(a int32, b []int32, incr []int32) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVI64(a int64, b []int64, incr []int64) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVU(a uint, b []uint, incr []uint) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVU8(a uint8, b []uint8, incr []uint8) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVU16(a uint16, b []uint16, incr []uint16) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVU32(a uint32, b []uint32, incr []uint32) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVU64(a uint64, b []uint64, incr []uint64) {
	for i := range incr {
		incr[i] += a % b[i]
	}
}

func ModIncrSVF32(a float32, b []float32, incr []float32) {
	for i := range incr {
		incr[i] += math32.Mod(a, b[i])
	}
}

func ModIncrSVF64(a float64, b []float64, incr []float64) {
	for i := range incr {
		incr[i] += math.Mod(a, b[i])
	}
}

func AddIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVC64(a complex64, b []complex64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVC128(a complex128, b []complex128, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func AddIterSVStr(a string, b []string, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a + b[i]
		}
	}
	return
}

func SubIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVC64(a complex64, b []complex64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func SubIterSVC128(a complex128, b []complex128, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a - b[i]
		}
	}
	return
}

func MulIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVC64(a complex64, b []complex64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func MulIterSVC128(a complex128, b []complex128, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a * b[i]
		}
	}
	return
}

func DivIterSVI(a int, b []int, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b[i] == 0 {
				errs = append(errs, i)
				b[i] = 0
				continue
			}
			b[i] = a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a / b[i]
		}
	}
	return
}

func DivIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a / b[i]
		}
	}
	return
}

func DivIterSVC64(a complex64, b []complex64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a / b[i]
		}
	}
	return
}

func DivIterSVC128(a complex128, b []complex128, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a / b[i]
		}
	}
	return
}

func PowIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = math32.Pow(a, b[i])
		}
	}
	return
}

func PowIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = math.Pow(a, b[i])
		}
	}
	return
}

func PowIterSVC64(a complex64, b []complex64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = complex64(cmplx.Pow(complex128(a), complex128(b[i])))
		}
	}
	return
}

func PowIterSVC128(a complex128, b []complex128, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = cmplx.Pow(a, b[i])
		}
	}
	return
}

func ModIterSVI(a int, b []int, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVI8(a int8, b []int8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVI16(a int16, b []int16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVI32(a int32, b []int32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVI64(a int64, b []int64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVU(a uint, b []uint, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVU8(a uint8, b []uint8, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVU16(a uint16, b []uint16, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVU32(a uint32, b []uint32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVU64(a uint64, b []uint64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = a % b[i]
		}
	}
	return
}

func ModIterSVF32(a float32, b []float32, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = math32.Mod(a, b[i])
		}
	}
	return
}

func ModIterSVF64(a float64, b []float64, bit Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			b[i] = math.Mod(a, b[i])
		}
	}
	return
}

func AddIterIncrSVI(a int, b []int, incr []int, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVI8(a int8, b []int8, incr []int8, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVI16(a int16, b []int16, incr []int16, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVI32(a int32, b []int32, incr []int32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVI64(a int64, b []int64, incr []int64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVU(a uint, b []uint, incr []uint, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVU8(a uint8, b []uint8, incr []uint8, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVU16(a uint16, b []uint16, incr []uint16, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVU32(a uint32, b []uint32, incr []uint32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVU64(a uint64, b []uint64, incr []uint64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVF32(a float32, b []float32, incr []float32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVF64(a float64, b []float64, incr []float64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVC64(a complex64, b []complex64, incr []complex64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVC128(a complex128, b []complex128, incr []complex128, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func AddIterIncrSVStr(a string, b []string, incr []string, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a + b[i]
		}
	}
	return
}

func SubIterIncrSVI(a int, b []int, incr []int, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVI8(a int8, b []int8, incr []int8, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVI16(a int16, b []int16, incr []int16, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVI32(a int32, b []int32, incr []int32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVI64(a int64, b []int64, incr []int64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVU(a uint, b []uint, incr []uint, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVU8(a uint8, b []uint8, incr []uint8, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVU16(a uint16, b []uint16, incr []uint16, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVU32(a uint32, b []uint32, incr []uint32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVU64(a uint64, b []uint64, incr []uint64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVF32(a float32, b []float32, incr []float32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVF64(a float64, b []float64, incr []float64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVC64(a complex64, b []complex64, incr []complex64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func SubIterIncrSVC128(a complex128, b []complex128, incr []complex128, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a - b[i]
		}
	}
	return
}

func MulIterIncrSVI(a int, b []int, incr []int, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVI8(a int8, b []int8, incr []int8, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVI16(a int16, b []int16, incr []int16, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVI32(a int32, b []int32, incr []int32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVI64(a int64, b []int64, incr []int64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVU(a uint, b []uint, incr []uint, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVU8(a uint8, b []uint8, incr []uint8, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVU16(a uint16, b []uint16, incr []uint16, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVU32(a uint32, b []uint32, incr []uint32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVU64(a uint64, b []uint64, incr []uint64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVF32(a float32, b []float32, incr []float32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVF64(a float64, b []float64, incr []float64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVC64(a complex64, b []complex64, incr []complex64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func MulIterIncrSVC128(a complex128, b []complex128, incr []complex128, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a * b[i]
		}
	}
	return
}

func DivIterIncrSVI(a int, b []int, incr []int, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVI8(a int8, b []int8, incr []int8, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVI16(a int16, b []int16, incr []int16, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVI32(a int32, b []int32, incr []int32, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVI64(a int64, b []int64, incr []int64, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVU(a uint, b []uint, incr []uint, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVU8(a uint8, b []uint8, incr []uint8, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVU16(a uint16, b []uint16, incr []uint16, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVU32(a uint32, b []uint32, incr []uint32, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVU64(a uint64, b []uint64, incr []uint64, bit Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b[i] == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a / b[i]
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrSVF32(a float32, b []float32, incr []float32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a / b[i]
		}
	}
	return
}

func DivIterIncrSVF64(a float64, b []float64, incr []float64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a / b[i]
		}
	}
	return
}

func DivIterIncrSVC64(a complex64, b []complex64, incr []complex64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a / b[i]
		}
	}
	return
}

func DivIterIncrSVC128(a complex128, b []complex128, incr []complex128, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a / b[i]
		}
	}
	return
}

func PowIterIncrSVF32(a float32, b []float32, incr []float32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += math32.Pow(a, b[i])
		}
	}
	return
}

func PowIterIncrSVF64(a float64, b []float64, incr []float64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += math.Pow(a, b[i])
		}
	}
	return
}

func PowIterIncrSVC64(a complex64, b []complex64, incr []complex64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += complex64(cmplx.Pow(complex128(a), complex128(b[i])))
		}
	}
	return
}

func PowIterIncrSVC128(a complex128, b []complex128, incr []complex128, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += cmplx.Pow(a, b[i])
		}
	}
	return
}

func ModIterIncrSVI(a int, b []int, incr []int, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVI8(a int8, b []int8, incr []int8, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVI16(a int16, b []int16, incr []int16, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVI32(a int32, b []int32, incr []int32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVI64(a int64, b []int64, incr []int64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVU(a uint, b []uint, incr []uint, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVU8(a uint8, b []uint8, incr []uint8, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVU16(a uint16, b []uint16, incr []uint16, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVU32(a uint32, b []uint32, incr []uint32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVU64(a uint64, b []uint64, incr []uint64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a % b[i]
		}
	}
	return
}

func ModIterIncrSVF32(a float32, b []float32, incr []float32, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += math32.Mod(a, b[i])
		}
	}
	return
}

func ModIterIncrSVF64(a float64, b []float64, incr []float64, bit Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = bit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += math.Mod(a, b[i])
		}
	}
	return
}

func AddVSI(a []int, b int) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSI8(a []int8, b int8) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSI16(a []int16, b int16) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSI32(a []int32, b int32) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSI64(a []int64, b int64) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSU(a []uint, b uint) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSU8(a []uint8, b uint8) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSU16(a []uint16, b uint16) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSU32(a []uint32, b uint32) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSU64(a []uint64, b uint64) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSF32(a []float32, b float32) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSF64(a []float64, b float64) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSC64(a []complex64, b complex64) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSC128(a []complex128, b complex128) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func AddVSStr(a []string, b string) {
	for i := range a {
		a[i] = a[i] + b
	}
}

func SubVSI(a []int, b int) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSI8(a []int8, b int8) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSI16(a []int16, b int16) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSI32(a []int32, b int32) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSI64(a []int64, b int64) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSU(a []uint, b uint) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSU8(a []uint8, b uint8) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSU16(a []uint16, b uint16) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSU32(a []uint32, b uint32) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSU64(a []uint64, b uint64) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSF32(a []float32, b float32) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSF64(a []float64, b float64) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSC64(a []complex64, b complex64) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func SubVSC128(a []complex128, b complex128) {
	for i := range a {
		a[i] = a[i] - b
	}
}

func MulVSI(a []int, b int) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSI8(a []int8, b int8) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSI16(a []int16, b int16) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSI32(a []int32, b int32) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSI64(a []int64, b int64) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSU(a []uint, b uint) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSU8(a []uint8, b uint8) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSU16(a []uint16, b uint16) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSU32(a []uint32, b uint32) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSU64(a []uint64, b uint64) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSF32(a []float32, b float32) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSF64(a []float64, b float64) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSC64(a []complex64, b complex64) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func MulVSC128(a []complex128, b complex128) {
	for i := range a {
		a[i] = a[i] * b
	}
}

func DivVSI(a []int, b int) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSI8(a []int8, b int8) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSI16(a []int16, b int16) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSI32(a []int32, b int32) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSI64(a []int64, b int64) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSU(a []uint, b uint) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSU8(a []uint8, b uint8) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSU16(a []uint16, b uint16) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSU32(a []uint32, b uint32) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSU64(a []uint64, b uint64) (err error) {
	var errs errorIndices
	for i := range a {
		if b == 0 {
			errs = append(errs, i)
			a[i] = 0
			continue
		}
		a[i] = a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivVSF32(a []float32, b float32) {
	for i := range a {
		a[i] = a[i] / b
	}
}

func DivVSF64(a []float64, b float64) {
	for i := range a {
		a[i] = a[i] / b
	}
}

func DivVSC64(a []complex64, b complex64) {
	for i := range a {
		a[i] = a[i] / b
	}
}

func DivVSC128(a []complex128, b complex128) {
	for i := range a {
		a[i] = a[i] / b
	}
}

func PowVSF32(a []float32, b float32) {
	for i := range a {
		a[i] = math32.Pow(a[i], b)
	}
}

func PowVSF64(a []float64, b float64) {
	for i := range a {
		a[i] = math.Pow(a[i], b)
	}
}

func PowVSC64(a []complex64, b complex64) {
	for i := range a {
		a[i] = complex64(cmplx.Pow(complex128(a[i]), complex128(b)))
	}
}

func PowVSC128(a []complex128, b complex128) {
	for i := range a {
		a[i] = cmplx.Pow(a[i], b)
	}
}

func ModVSI(a []int, b int) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSI8(a []int8, b int8) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSI16(a []int16, b int16) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSI32(a []int32, b int32) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSI64(a []int64, b int64) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSU(a []uint, b uint) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSU8(a []uint8, b uint8) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSU16(a []uint16, b uint16) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSU32(a []uint32, b uint32) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSU64(a []uint64, b uint64) {
	for i := range a {
		a[i] = a[i] % b
	}
}

func ModVSF32(a []float32, b float32) {
	for i := range a {
		a[i] = math32.Mod(a[i], b)
	}
}

func ModVSF64(a []float64, b float64) {
	for i := range a {
		a[i] = math.Mod(a[i], b)
	}
}

func AddIncrVSI(a []int, b int, incr []int) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSI8(a []int8, b int8, incr []int8) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSI16(a []int16, b int16, incr []int16) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSI32(a []int32, b int32, incr []int32) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSI64(a []int64, b int64, incr []int64) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSU(a []uint, b uint, incr []uint) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSU8(a []uint8, b uint8, incr []uint8) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSU16(a []uint16, b uint16, incr []uint16) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSU32(a []uint32, b uint32, incr []uint32) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSU64(a []uint64, b uint64, incr []uint64) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSF32(a []float32, b float32, incr []float32) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSF64(a []float64, b float64, incr []float64) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSC64(a []complex64, b complex64, incr []complex64) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSC128(a []complex128, b complex128, incr []complex128) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func AddIncrVSStr(a []string, b string, incr []string) {
	for i := range incr {
		incr[i] += a[i] + b
	}
}

func SubIncrVSI(a []int, b int, incr []int) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSI8(a []int8, b int8, incr []int8) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSI16(a []int16, b int16, incr []int16) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSI32(a []int32, b int32, incr []int32) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSI64(a []int64, b int64, incr []int64) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSU(a []uint, b uint, incr []uint) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSU8(a []uint8, b uint8, incr []uint8) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSU16(a []uint16, b uint16, incr []uint16) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSU32(a []uint32, b uint32, incr []uint32) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSU64(a []uint64, b uint64, incr []uint64) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSF32(a []float32, b float32, incr []float32) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSF64(a []float64, b float64, incr []float64) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSC64(a []complex64, b complex64, incr []complex64) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func SubIncrVSC128(a []complex128, b complex128, incr []complex128) {
	for i := range incr {
		incr[i] += a[i] - b
	}
}

func MulIncrVSI(a []int, b int, incr []int) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSI8(a []int8, b int8, incr []int8) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSI16(a []int16, b int16, incr []int16) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSI32(a []int32, b int32, incr []int32) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSI64(a []int64, b int64, incr []int64) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSU(a []uint, b uint, incr []uint) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSU8(a []uint8, b uint8, incr []uint8) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSU16(a []uint16, b uint16, incr []uint16) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSU32(a []uint32, b uint32, incr []uint32) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSU64(a []uint64, b uint64, incr []uint64) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSF32(a []float32, b float32, incr []float32) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSF64(a []float64, b float64, incr []float64) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSC64(a []complex64, b complex64, incr []complex64) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func MulIncrVSC128(a []complex128, b complex128, incr []complex128) {
	for i := range incr {
		incr[i] += a[i] * b
	}
}

func DivIncrVSI(a []int, b int, incr []int) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSI8(a []int8, b int8, incr []int8) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSI16(a []int16, b int16, incr []int16) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSI32(a []int32, b int32, incr []int32) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSI64(a []int64, b int64, incr []int64) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSU(a []uint, b uint, incr []uint) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSU8(a []uint8, b uint8, incr []uint8) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSU16(a []uint16, b uint16, incr []uint16) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSU32(a []uint32, b uint32, incr []uint32) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSU64(a []uint64, b uint64, incr []uint64) (err error) {
	var errs errorIndices
	for i := range incr {
		if b == 0 {
			errs = append(errs, i)
			incr[i] = 0
			continue
		}
		incr[i] += a[i] / b
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIncrVSF32(a []float32, b float32, incr []float32) {
	for i := range incr {
		incr[i] += a[i] / b
	}
}

func DivIncrVSF64(a []float64, b float64, incr []float64) {
	for i := range incr {
		incr[i] += a[i] / b
	}
}

func DivIncrVSC64(a []complex64, b complex64, incr []complex64) {
	for i := range incr {
		incr[i] += a[i] / b
	}
}

func DivIncrVSC128(a []complex128, b complex128, incr []complex128) {
	for i := range incr {
		incr[i] += a[i] / b
	}
}

func PowIncrVSF32(a []float32, b float32, incr []float32) {
	for i := range incr {
		incr[i] += math32.Pow(a[i], b)
	}
}

func PowIncrVSF64(a []float64, b float64, incr []float64) {
	for i := range incr {
		incr[i] += math.Pow(a[i], b)
	}
}

func PowIncrVSC64(a []complex64, b complex64, incr []complex64) {
	for i := range incr {
		incr[i] += complex64(cmplx.Pow(complex128(a[i]), complex128(b)))
	}
}

func PowIncrVSC128(a []complex128, b complex128, incr []complex128) {
	for i := range incr {
		incr[i] += cmplx.Pow(a[i], b)
	}
}

func ModIncrVSI(a []int, b int, incr []int) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSI8(a []int8, b int8, incr []int8) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSI16(a []int16, b int16, incr []int16) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSI32(a []int32, b int32, incr []int32) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSI64(a []int64, b int64, incr []int64) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSU(a []uint, b uint, incr []uint) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSU8(a []uint8, b uint8, incr []uint8) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSU16(a []uint16, b uint16, incr []uint16) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSU32(a []uint32, b uint32, incr []uint32) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSU64(a []uint64, b uint64, incr []uint64) {
	for i := range incr {
		incr[i] += a[i] % b
	}
}

func ModIncrVSF32(a []float32, b float32, incr []float32) {
	for i := range incr {
		incr[i] += math32.Mod(a[i], b)
	}
}

func ModIncrVSF64(a []float64, b float64, incr []float64) {
	for i := range incr {
		incr[i] += math.Mod(a[i], b)
	}
}

func AddIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSC64(a []complex64, b complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSC128(a []complex128, b complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func AddIterVSStr(a []string, b string, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] + b
		}
	}
	return
}

func SubIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSC64(a []complex64, b complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func SubIterVSC128(a []complex128, b complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] - b
		}
	}
	return
}

func MulIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSC64(a []complex64, b complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func MulIterVSC128(a []complex128, b complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] * b
		}
	}
	return
}

func DivIterVSI(a []int, b int, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var errs errorIndices
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			if b == 0 {
				errs = append(errs, i)
				a[i] = 0
				continue
			}
			a[i] = a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] / b
		}
	}
	return
}

func DivIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] / b
		}
	}
	return
}

func DivIterVSC64(a []complex64, b complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] / b
		}
	}
	return
}

func DivIterVSC128(a []complex128, b complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] / b
		}
	}
	return
}

func PowIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Pow(a[i], b)
		}
	}
	return
}

func PowIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Pow(a[i], b)
		}
	}
	return
}

func PowIterVSC64(a []complex64, b complex64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = complex64(cmplx.Pow(complex128(a[i]), complex128(b)))
		}
	}
	return
}

func PowIterVSC128(a []complex128, b complex128, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = cmplx.Pow(a[i], b)
		}
	}
	return
}

func ModIterVSI(a []int, b int, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSI8(a []int8, b int8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSI16(a []int16, b int16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSI32(a []int32, b int32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSI64(a []int64, b int64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSU(a []uint, b uint, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSU8(a []uint8, b uint8, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSU16(a []uint16, b uint16, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSU32(a []uint32, b uint32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSU64(a []uint64, b uint64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = a[i] % b
		}
	}
	return
}

func ModIterVSF32(a []float32, b float32, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math32.Mod(a[i], b)
		}
	}
	return
}

func ModIterVSF64(a []float64, b float64, ait Iterator) (err error) {
	var i int
	var validi bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi {
			a[i] = math.Mod(a[i], b)
		}
	}
	return
}

func AddIterIncrVSI(a []int, b int, incr []int, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSI8(a []int8, b int8, incr []int8, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSI16(a []int16, b int16, incr []int16, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSI32(a []int32, b int32, incr []int32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSI64(a []int64, b int64, incr []int64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSU(a []uint, b uint, incr []uint, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSU8(a []uint8, b uint8, incr []uint8, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSU16(a []uint16, b uint16, incr []uint16, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSU32(a []uint32, b uint32, incr []uint32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSU64(a []uint64, b uint64, incr []uint64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSF32(a []float32, b float32, incr []float32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSF64(a []float64, b float64, incr []float64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSC64(a []complex64, b complex64, incr []complex64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSC128(a []complex128, b complex128, incr []complex128, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func AddIterIncrVSStr(a []string, b string, incr []string, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] + b
		}
	}
	return
}

func SubIterIncrVSI(a []int, b int, incr []int, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSI8(a []int8, b int8, incr []int8, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSI16(a []int16, b int16, incr []int16, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSI32(a []int32, b int32, incr []int32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSI64(a []int64, b int64, incr []int64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSU(a []uint, b uint, incr []uint, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSU8(a []uint8, b uint8, incr []uint8, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSU16(a []uint16, b uint16, incr []uint16, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSU32(a []uint32, b uint32, incr []uint32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSU64(a []uint64, b uint64, incr []uint64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSF32(a []float32, b float32, incr []float32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSF64(a []float64, b float64, incr []float64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSC64(a []complex64, b complex64, incr []complex64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func SubIterIncrVSC128(a []complex128, b complex128, incr []complex128, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] - b
		}
	}
	return
}

func MulIterIncrVSI(a []int, b int, incr []int, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSI8(a []int8, b int8, incr []int8, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSI16(a []int16, b int16, incr []int16, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSI32(a []int32, b int32, incr []int32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSI64(a []int64, b int64, incr []int64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSU(a []uint, b uint, incr []uint, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSU8(a []uint8, b uint8, incr []uint8, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSU16(a []uint16, b uint16, incr []uint16, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSU32(a []uint32, b uint32, incr []uint32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSU64(a []uint64, b uint64, incr []uint64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSF32(a []float32, b float32, incr []float32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSF64(a []float64, b float64, incr []float64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSC64(a []complex64, b complex64, incr []complex64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func MulIterIncrVSC128(a []complex128, b complex128, incr []complex128, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] * b
		}
	}
	return
}

func DivIterIncrVSI(a []int, b int, incr []int, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSI8(a []int8, b int8, incr []int8, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSI16(a []int16, b int16, incr []int16, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSI32(a []int32, b int32, incr []int32, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSI64(a []int64, b int64, incr []int64, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSU(a []uint, b uint, incr []uint, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSU8(a []uint8, b uint8, incr []uint8, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSU16(a []uint16, b uint16, incr []uint16, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSU32(a []uint32, b uint32, incr []uint32, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSU64(a []uint64, b uint64, incr []uint64, ait Iterator, iit Iterator) (err error) {
	var errs errorIndices
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			if b == 0 {
				errs = append(errs, i)
				incr[i] = 0
				continue
			}
			incr[k] += a[i] / b
		}
	}
	if err != nil {
		return
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func DivIterIncrVSF32(a []float32, b float32, incr []float32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] / b
		}
	}
	return
}

func DivIterIncrVSF64(a []float64, b float64, incr []float64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] / b
		}
	}
	return
}

func DivIterIncrVSC64(a []complex64, b complex64, incr []complex64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] / b
		}
	}
	return
}

func DivIterIncrVSC128(a []complex128, b complex128, incr []complex128, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] / b
		}
	}
	return
}

func PowIterIncrVSF32(a []float32, b float32, incr []float32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += math32.Pow(a[i], b)
		}
	}
	return
}

func PowIterIncrVSF64(a []float64, b float64, incr []float64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += math.Pow(a[i], b)
		}
	}
	return
}

func PowIterIncrVSC64(a []complex64, b complex64, incr []complex64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += complex64(cmplx.Pow(complex128(a[i]), complex128(b)))
		}
	}
	return
}

func PowIterIncrVSC128(a []complex128, b complex128, incr []complex128, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += cmplx.Pow(a[i], b)
		}
	}
	return
}

func ModIterIncrVSI(a []int, b int, incr []int, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSI8(a []int8, b int8, incr []int8, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSI16(a []int16, b int16, incr []int16, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSI32(a []int32, b int32, incr []int32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSI64(a []int64, b int64, incr []int64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSU(a []uint, b uint, incr []uint, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSU8(a []uint8, b uint8, incr []uint8, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSU16(a []uint16, b uint16, incr []uint16, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSU32(a []uint32, b uint32, incr []uint32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSU64(a []uint64, b uint64, incr []uint64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += a[i] % b
		}
	}
	return
}

func ModIterIncrVSF32(a []float32, b float32, incr []float32, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += math32.Mod(a[i], b)
		}
	}
	return
}

func ModIterIncrVSF64(a []float64, b float64, incr []float64, ait Iterator, iit Iterator) (err error) {
	var i, k int
	var validi, validk bool
	for {
		if i, validi, err = ait.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if k, validk, err = iit.NextValidity(); err != nil {
			err = handleNoOp(err)
			break
		}
		if validi && validk {
			incr[k] += math.Mod(a[i], b)
		}
	}
	return
}
