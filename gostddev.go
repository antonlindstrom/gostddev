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
	var d float64
	for i, v := range a {
		d = v - meanVal
		set[i] = d * d
	}
	return set
}

func StdDev(a []float64) float64 {
	return math.Sqrt(Mean(DiffSqrtMean(a)))
}
