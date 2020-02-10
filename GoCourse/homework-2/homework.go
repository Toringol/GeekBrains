package main

import "fmt"

// 1) Написать функцию, которая определяет, четное ли число.
func isEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

// 2) Написать функцию, которая определяет, делится ли число без остатка на 3.
func isDivideOnThree(number int) bool {
	if number%3 == 0 {
		return true
	}
	return false
}

// 3) Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
func fibonacciToHundred() {
	var fibNumber uint = 0
	var prevFibNumber uint = 1
	var prevPrevFibNumber uint = 0

	for i := 0; i < 100; i++ {
		if i == 0 {
			fibNumber = 0
		} else if i == 1 {
			fibNumber = 1
		} else {
			fibNumber = prevFibNumber + prevPrevFibNumber
			prevPrevFibNumber = prevFibNumber
			prevFibNumber = fibNumber
		}
		fmt.Println(i, " ", fibNumber)
	}
}

// 4) Заполнить массив из 100 элементов различными простыми числами.
const (
	sizeResultArray = 100
	sizeTempArray   = 1000 // Берем временнй массив заведомо больше
)

func generateHundredElementsArrayPrimeNumbers() [sizeResultArray]int {
	var arrPrimeNumbers = [sizeResultArray]int{}
	var arrСonsecutiveNumbers = [sizeTempArray]int{}

	// Заполняем временный массив
	for i := 0; i < sizeTempArray; i++ {
		arrСonsecutiveNumbers[i] = i
	}

	// Идем по всем числам, начиная от 2, при этом следим за переполнением целевого массива
	var iterOverflowArray int = 0
	for p := 2; p < sizeTempArray && iterOverflowArray < sizeResultArray; p++ {
		if arrСonsecutiveNumbers[p] != 0 {

			// Записываем первое не зачеркнутое число в списке
			arrPrimeNumbers[iterOverflowArray] = arrСonsecutiveNumbers[p]
			iterOverflowArray++

			// Вычеркиваем все кратные числа
			for j := p * p; j < sizeTempArray; j += p {
				arrСonsecutiveNumbers[j] = 0
			}
		}
	}

	return arrPrimeNumbers
}

func main() {
	fibonacciToHundred()
	fmt.Println(generateHundredElementsArrayPrimeNumbers())
}
