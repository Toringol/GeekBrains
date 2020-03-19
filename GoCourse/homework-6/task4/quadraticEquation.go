package task4

import "math"

func SolutionQuadraticEquation(a, b, c float64) []float64 {
	resultRoots := []float64{}

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant > 0 {
		firstRoot := (-b + math.Sqrt(discriminant)) / (2 * a)
		secondRoot := (-b - math.Sqrt(discriminant)) / (2 * a)

		resultRoots = append(resultRoots, firstRoot)
		resultRoots = append(resultRoots, secondRoot)
	} else if discriminant == 0 {
		root := -b / (2 * a)

		resultRoots = append(resultRoots, root)
	}

	return resultRoots
}

func SolutionQuadraticEquationFaster(a, b, c float64) []float64 {
	resultRoots := []float64{}

	discriminant := b*b - 2*2*a*c

	if discriminant > 0 {
		firstRoot := (-b + math.Sqrt(discriminant)) / (2 * a)
		secondRoot := (-b - math.Sqrt(discriminant)) / (2 * a)

		resultRoots = append(resultRoots, firstRoot)
		resultRoots = append(resultRoots, secondRoot)
	} else if discriminant == 0 {
		root := -b / (2 * a)

		resultRoots = append(resultRoots, root)
	}

	return resultRoots
}
