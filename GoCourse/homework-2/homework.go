package main

import (
	"fmt"
	"math/big"
)

// 1) и 2) Определяет делится ли число на другое число без остатка
func isDivide(number int, div int) bool {
	return number%div == 0
}

// 3) Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
func fibonacciToHundred() {
	fibNumber := big.NewInt(1)
	prevFibNumber := big.NewInt(1)

	for i := 2; i <= 100; i++ {
		fibNumber.Add(fibNumber, prevFibNumber)
		fibNumber, prevFibNumber = prevFibNumber, fibNumber
		fmt.Println(i, " ", fibNumber)
	}
}

// 4) Заполнить массив из 100 элементов различными простыми числами.
const (
	sizeResultArray = 100
	sizeTempArray   = 1000 // Берем временнй массив заведомо больше
)

func generateHundredElementsArrayPrimeNumbers() [sizeResultArray]int {
	var arrPrimeNumbers [sizeResultArray]int
	var arrСonsecutiveNumbers [sizeTempArray]int

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
