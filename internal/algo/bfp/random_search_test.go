package bfp_test

import (
	"math"
	"testing"

	bfp "github.com/alelince/metaheuristics-playground/internal/algo/bfp"
)

func TestRandomSearchFindMinimum(t *testing.T) {
	const tol float64 = 0.5

	randomSearch := bfp.NewRandomSearch(bfp.UNIFORM, int(1e6), 4)
	f := &function{}

	var tests = []struct {
		name     string
		function bfp.IFunction
		xMin     float64
		xMax     float64
		minimum  []float64
	}{
		{"x^2 from -5 to 5", f, -5, 5, []float64{0., 0.}},
		{"x^2 from 1 to 20", f, 1, 20, []float64{1., 1.}},
		{"x^2 from -10 to -1", f, -10, -1, []float64{-1., -1.}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minVec, err := randomSearch.FindMinimum(tt.function, tt.xMin, tt.xMax, len(tt.minimum))
			if err != nil {
				t.Errorf("test %s: computation error: %v", tt.name, err)
			}

			if math.Abs(tt.minimum[0]-minVec[0]) > tol || math.Abs(tt.minimum[1]-minVec[1]) > tol {
				t.Errorf("test %s: computed minimum vector too different: %v", tt.name, minVec)
			}
		})
	}
}
