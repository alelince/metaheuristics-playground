package bfp

import (
	"math"
	"math/rand/v2"
)

type IFunction interface {
	Compute(x []float64) (float64, error)
}

func UniformRandomVector(min, max float64, n int) []float64 {
	vector := make([]float64, n)

	for i := 0; i < n; i++ {
		vector[i] = min + (max-min)*rand.Float64()
	}

	return vector
}

func FindMinimumAmongMinima(f IFunction, minVectors [][]float64) (int, []float64) {
	minIndex := -1
	minValue := math.MaxFloat64

	// search the minimum among the minima obtained by the goroutines
	for i := 0; i < len(minVectors); i++ {
		auxValue, err := f.Compute(minVectors[i])
		if err != nil {
			continue // ignore minimum value in case of computation error
		}

		if auxValue < minValue {
			minIndex = i
			minValue = auxValue
		}
	}

	if minIndex == -1 {
		return -1, nil
	}

	return minIndex, minVectors[minIndex]
}
