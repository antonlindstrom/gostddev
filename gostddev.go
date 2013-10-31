package gostddev

import "math"

func Mean(a []float64) float64 {
	return Sum(a) / float64(len(a))
}

func Sum(a []float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}

func DiffSqrtMean(a []float64) []float64 {
	set := make([]float64, len(a))
	meanVal := Mean(a)

	for i, v := range a {
		set[i] = math.Pow(v-meanVal, 2)
	}
	return set
}

func StdDev(a []float64) float64 {
	return math.Sqrt(Mean(DiffSqrtMean(a)))
}
