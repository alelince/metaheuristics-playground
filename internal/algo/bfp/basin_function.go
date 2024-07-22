package bfp

import (
	"fmt"
	"math"
)

const (
	MAX_N = 10
)

type BasinFunction struct {
	A, H, K float64
}

func NewBasinFunction(a, h, k float64) *BasinFunction {
	return &BasinFunction{A: a, H: h, K: k}
}

func (b *BasinFunction) Compute(x []float64) (float64, error) {
	// sanity check
	if len(x) < 1 || len(x) > MAX_N {
		return 0, fmt.Errorf("invalid input size in basin function: %v", x)
	}

	var result float64 = 0.

	for _, xi := range x {
		result += b.A*math.Pow((xi-b.H), 2) + b.K
	}

	return result, nil
}
