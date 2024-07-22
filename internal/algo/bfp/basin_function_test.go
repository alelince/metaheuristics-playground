package bfp_test

import (
	"math"
	"testing"

	bfp "github.com/alelince/metaheuristics-playground/internal/algo/bfp"
)

func TestBasinFunctionCompute(t *testing.T) {
	bf := bfp.NewBasinFunction(0.5, 2., -5.)
	tol := 1e-3

	var tests = []struct {
		name  string
		x     []float64
		value float64
	}{
		{"(0, 0) gives -6", []float64{0., 0.}, -6.},
		{"(2, 2) gives -10", []float64{2., 2.}, -10.},
		{"(3, 3) gives -9", []float64{3., 3.}, -9.},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := bf.Compute(tt.x)
			if err != nil {
				t.Errorf("test %s: computation error: %v", tt.name, err)
			}

			if math.Abs(val-tt.value) > tol {
				t.Errorf("test %s: computed value too different: %f", tt.name, val)
			}
		})
	}
}
