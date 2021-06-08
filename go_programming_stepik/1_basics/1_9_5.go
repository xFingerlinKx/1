/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".
*/
package main

import "fmt"

func main() {
	var inputInt int
	fmt.Scan(&inputInt)
	firstNum := inputInt / 100
	secondNum := inputInt / 10 % 10
	thirdNum := inputInt % 10

	switch {
	case firstNum != secondNum && secondNum != thirdNum && thirdNum != firstNum:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
