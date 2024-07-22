package bfp

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

const (
	UNIFORM = iota
)

type RandomSearch struct {
	ProbabilityDistribution int
	MaxIterations           int
	NumGoRoutines           int
}

func NewRandomSearch(probDistro, maxIterations, numGoRoutines int) *RandomSearch {
	return &RandomSearch{
		ProbabilityDistribution: probDistro,
		MaxIterations:           maxIterations,
		NumGoRoutines:           numGoRoutines,
	}
}

func (r *RandomSearch) FindMinimum(f IFunction, min, max float64, n int) ([]float64, error) {
	var wg sync.WaitGroup
	wg.Add(r.NumGoRoutines)

	var iterationCounter atomic.Int64
	minVectors := make([][]float64, r.NumGoRoutines)

	// distribute workload in different goroutines
	for i := 0; i < r.NumGoRoutines; i++ {
		go func() {
			defer wg.Done()

			minValue := math.MaxFloat64

			for {
				if iterationCounter.Add(1) > int64(r.MaxIterations) {
					return
				}

				// random search algorithm
				candidateVector := r.randomSample(min, max, n)
				candidateValue, err := f.Compute(candidateVector)
				if err != nil {
					continue // ignore sample in case of computation error
				}

				if candidateValue < minValue {
					minVectors[i] = candidateVector
					minValue = candidateValue
				}
			}
		}()
	}

	wg.Wait()

	_, minVector := FindMinimumAmongMinima(f, minVectors)

	if minVector == nil {
		return nil, fmt.Errorf(
			"error in random search with min = %f, max = %f, n = %d ", min, max, n,
		)
	}

	return minVector, nil
}

func (r *RandomSearch) randomSample(min, max float64, n int) []float64 {
	var vector []float64

	switch r.ProbabilityDistribution {
	case UNIFORM:
		fallthrough
	default:
		vector = UniformRandomVector(min, max, n)
	}

	return vector
}
