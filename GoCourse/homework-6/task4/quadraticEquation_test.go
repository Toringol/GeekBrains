package task4

import (
	"reflect"
	"testing"
)

type testpair struct {
	input    []float64
	expected []float64
}

var tests = []testpair{
	{[]float64{2, 5, 3}, []float64{-1, -1.5}},
	{[]float64{2, 5, 2}, []float64{-0.5, -2}},
	{[]float64{1, 1, 1}, []float64{}},
	{[]float64{1, 2, 1}, []float64{-1}},
}

func TestQuadraticEquation(t *testing.T) {
	for _, pair := range tests {
		v := SolutionQuadraticEquation(pair.input[0], pair.input[1], pair.input[2])
		if !reflect.DeepEqual(v, pair.expected) {
			t.Error(
				"For", pair.input,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func TestQuadraticEquationFaster(t *testing.T) {
	for _, pair := range tests {
		v := SolutionQuadraticEquationFaster(pair.input[0], pair.input[1], pair.input[2])
		if !reflect.DeepEqual(v, pair.expected) {
			t.Error(
				"For", pair.input,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func BenchmarkQuadraticEquationWithTwoRoots(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SolutionQuadraticEquation(tests[0].input[0], tests[0].input[1], tests[0].input[2])
	}
}

func BenchmarkQuadraticEquationWithTwoRootsFaster(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SolutionQuadraticEquationFaster(tests[0].input[0], tests[0].input[1], tests[0].input[2])
	}
}

func BenchmarkQuadraticEquationWithOneRoot(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SolutionQuadraticEquation(tests[3].input[0], tests[3].input[1], tests[3].input[2])
	}
}

func BenchmarkQuadraticEquationWithOneRootFaster(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SolutionQuadraticEquationFaster(tests[3].input[0], tests[3].input[1], tests[3].input[2])
	}
}

func BenchmarkQuadraticEquationWithNoRoots(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SolutionQuadraticEquation(tests[2].input[0], tests[2].input[1], tests[2].input[2])
	}
}

func BenchmarkQuadraticEquationWithNoRootsFaster(t *testing.B) {
	for i := 0; i < t.N; i++ {
		SolutionQuadraticEquationFaster(tests[2].input[0], tests[2].input[1], tests[2].input[2])
	}
}
