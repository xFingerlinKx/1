/*
Дано натуральное число, выведите его последнюю цифру.
*/
package main

import (
	"fmt"
)

func main() {
	var inputInt int
	fmt.Scan(&inputInt)
	fmt.Println(inputInt % 10)
}
