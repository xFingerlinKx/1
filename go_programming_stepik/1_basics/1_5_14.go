/*
Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).
*/
package main

import "fmt"

func main() {
	var inputInt int
	fmt.Scan(&inputInt)
	fmt.Println(inputInt % 100 / 10)
}
