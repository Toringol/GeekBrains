package task1

import "testing"

type testpair struct {
	values   []float64
	expected float64
}

var testsAverage = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var testSum = []testpair{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
	{[]float64{6, 10}, 16},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range testsAverage {
		v := Average(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func TestSum(t *testing.T) {
	for _, pair := range testSum {
		v := Sum(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
