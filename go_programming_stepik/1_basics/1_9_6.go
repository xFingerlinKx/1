/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.
*/
package main

import "fmt"

func main() {
	var firstDigit, number int
	fmt.Scanln(&number)

	firstDigit = number
	for firstDigit >= 10 {
		firstDigit = firstDigit / 10
	}
	fmt.Println(firstDigit)
}
