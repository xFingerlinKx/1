package main

import "fmt"

/*
Напишите программу, которая меняет местами два числа (x := 1; y := 2; swap(&x, &y) должно дать x=2 и y=1).
*/

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x, y := 1, 2
	fmt.Println("First x, y -", x, y)
	swap(&x, &y)
	fmt.Println("Second x, y -", x, y)
}
