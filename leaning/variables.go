package main

import "fmt"

/*
Переменные в Go создаются с помощью ключевого слова var,
за которым следуют имя переменной (x), тип (string) и присваиваемое значение.
Тип переменной можно пропускать:
	var x = "Text"

Можно записать так:
	var x string
    x = "Hello World"
    fmt.Println(x)

	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	Если мы хотим присвоить значение переменной при её создании,
	то можем использовать сокращенную запись:
	x := 5
	Тип в данном случае указывать не обязательно,
	так как компилятор Go способен определить тип по литералу,
	которым мы инициализируем переменную.

	В общем, желательно всегда использовать краткий вариант написания.
*/
// константы
const constant = "Const"

// определение нескольких переменных
var (
	a = 5
	b = 10
	c = 15
)

func main() {
	var text string = "Hello, world!"
	fmt.Println(text)
	x := 123
	fmt.Println(x)
	fmt.Println(constant)
	fmt.Println(a, "+", b, "+", c)
}
