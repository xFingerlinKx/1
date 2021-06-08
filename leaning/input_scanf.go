package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var input float64
	// Scanf заполняет переменную input числом, введенным нами
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
