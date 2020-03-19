package main

import (
	"fmt"
)

func main() {

	const dollarRate = 64.1
	var rublesToConvert float64 = 0

	for {
		fmt.Print("Input the amount of rubles to convert: ")
		fmt.Scanln(&rublesToConvert)

		if rublesToConvert < 0 {
			fmt.Println("Please input positive value")
		} else {
			break
		}
	}

	resultAmount := rublesToConvert / dollarRate

	fmt.Println("Your amount in dollars: ", resultAmount)
}
