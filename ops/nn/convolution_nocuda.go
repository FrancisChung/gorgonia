// +build !cuda

package nnops

import (
	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func Conv2d(im, filter *G.Node, kernelShape tensor.Shape, pad, stride, dilation []int) (retVal *G.Node, err error) {
	return G.Conv2d(im, filter, kernelShape, pad, stride, dilation)
}
