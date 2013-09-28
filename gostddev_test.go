package gostddev

import "testing"

func TestMean(t *testing.T) {
	testSet := []float64{1, 2, 3, 4, 5}

	if Mean(testSet) != 3 {
		t.Error("expected 3 from set 1,2,3,4,5")
	}
}

func TestSum(t *testing.T) {
	testSet := []float64{1, 2, 3, 4, 5}

	if Sum(testSet) != 15 {
		t.Error("expected 15 from set 1,2,3,4,5")
	}
}

func TestStdDev(t *testing.T) {
	testSet := []float64{1, 2, 3, 4, 5}

	if StdDev(testSet) != 1.4142135623730951 {
		t.Error("expected 1.4142135623730951 from set 1,2,3,4,5")
	}
}
