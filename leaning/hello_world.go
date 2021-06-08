package main

// определение пакета - любая программа должна начинаться с этого

import "fmt"

// Пакет fmt (сокращение от format) - реализует форматирование для входных и выходных данных.

/*
Многострочный комментарий
*/

const name = "Artur"

func main() {
	fmt.Println("Hello, my name is", name)
	fmt.Println(len("Hello, my name is Artur"))
	fmt.Println(name[1])
	fmt.Println("1 + 1 =", 1.0+1.0)
	// && - И, || - ИЛИ, ! - НЕ
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
	fmt.Println(32132 * 42452)
}
