package main

import (
	"fmt"
	"math"
)

func main() {

	var firstCathet, secondCathet float64 = 0, 0

	for {
		fmt.Print("Input firstCather of right triangle: ")
		fmt.Scanln(&firstCathet)

		fmt.Print("Input secondCathet of right triangle: ")
		fmt.Scanln(&secondCathet)

		if firstCathet < 0 || secondCathet < 0 {
			fmt.Println("Please input positive values")
		} else {
			break
		}
	}

	triangleHypotenuse := math.Sqrt(firstCathet*firstCathet + secondCathet*secondCathet)
	trianglePerimeter := firstCathet + secondCathet + triangleHypotenuse
	triangleSquare := firstCathet * secondCathet / 2

	fmt.Println("Hypotenuse of right triangle: ", triangleHypotenuse)
	fmt.Println("Perimeter of right triangle: ", trianglePerimeter)
	fmt.Println("Square of right triangle: ", triangleSquare)

}
