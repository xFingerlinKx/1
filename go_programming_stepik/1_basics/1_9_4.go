/*
На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное",
если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль".
Выводить сообщение без кавычек.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	fmt.Scan(&number)
	if number == 0 {
		fmt.Println("Ноль")
	} else if math.Signbit(number) {
		fmt.Println("Число отрицательное")
	} else if !math.Signbit(number) {
		fmt.Println("Число положительное")
	}
}
