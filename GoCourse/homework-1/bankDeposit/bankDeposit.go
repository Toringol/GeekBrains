package main

import "fmt"

func main() {

	var depositAmount, depositPercentage, resultDeposit float64 = 0, 0, 0

	for {
		fmt.Print("Input depositAmount: ")
		fmt.Scanln(&depositAmount)

		fmt.Print("Input depositPercentage: ")
		fmt.Scanln(&depositPercentage)

		if depositAmount < 0 || depositPercentage < 0 {
			fmt.Println("Please input positive value")
		} else {
			break
		}
	}

	depositDuration := 5

	for i := 1; i <= depositDuration; i++ {
		depositAmount += depositAmount * depositPercentage / 100
	}

	resultDeposit = depositAmount

	fmt.Println("Deposit in 5 years: ", resultDeposit)

}
