package bfp_test

import (
	"math"
	"testing"

	bfp "github.com/alelince/metaheuristics-playground/internal/algo/bfp"
)

type function struct{}

// implement IFunction interface with an euclidean distance
func (f *function) Compute(x []float64) (float64, error) {
	aux := 0.

	for _, xi := range x {
		aux += math.Pow(xi, 2)
	}

	return math.Sqrt(aux), nil
}

func TestUniformRandomVector(t *testing.T) {
	const min float64 = -5.
	const max float64 = 5.
	const numSamples int = 10000
	const tol float64 = 0.1
	const mean float64 = (min + max) / 2.

	vector := bfp.UniformRandomVector(min, max, numSamples)
	calcMean := 0.

	for _, sample := range vector {
		if sample < min {
			t.Errorf("value below the minimum: %f", sample)
		} else if sample > max {
			t.Errorf("value above the maximum: %f", sample)
		}

		calcMean += sample
	}

	calcMean /= float64(numSamples)

	if calcMean > (tol+mean) || calcMean < (mean-tol) {
		t.Errorf("non uniform distribution! mean: %f", mean)
	}
}

func TestFindMinimumAmongMinima(t *testing.T) {
	f := &function{}

	vectors := make([][]float64, 5)
	vectors[0] = []float64{3., 3.}
	vectors[1] = []float64{1., 1.}
	vectors[2] = []float64{2., 2.}
	vectors[3] = []float64{4., 4.}
	vectors[4] = []float64{1., 2.}

	minIdx, minVec := bfp.FindMinimumAmongMinima(f, vectors)

	if minIdx != 1 {
		t.Errorf("vector index is wrong: %d", minIdx)
	}

	if minVec[0] != vectors[1][0] || minVec[1] != vectors[1][1] {
		t.Errorf("minimum vector is wrong: %v", minVec)
	}
}
