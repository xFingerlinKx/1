/*
Напишите программу, которая последовательно делает следующие операции с введённым числом:
Число умножается на 2;
Затем к числу прибавляется 100.
*/
package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)
	finalInput := input * 2 + 100
	fmt.Println(finalInput)
}
