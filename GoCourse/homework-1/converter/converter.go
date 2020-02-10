package main

import (
	"fmt"
)

func main() {

	const dollarRate = 64.1
	var rublesToConvert float64 = 0

	fmt.Print("Input the amount of rubles to convert: ")
	fmt.Scanln(&rublesToConvert)

	resultAmount := rublesToConvert / dollarRate

	fmt.Println("Your amount in dollars: ", resultAmount)
}
