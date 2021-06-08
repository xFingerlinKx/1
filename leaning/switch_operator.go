package main

import "fmt"

// Переключатель начинается с ключевого слова switch, за которым следует выражение
// (в нашем случае i) и серия возможных значений (case). Значение выражения по очереди
// сравнивается с выражениями, следующими после ключевого слова case. Если они оказываются равны,
// то выполняется действие, описанное после :.
func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println(i, "- Zero")
		case 1:
			fmt.Println(i, "- One")
		case 2:
			fmt.Println(i, "- Two")
		case 3:
			fmt.Println(i, "- Three")
		case 4:
			fmt.Println(i, "- Four")
		case 5:
			fmt.Println(i, "- Five")
		default:
			fmt.Println(i, "- Unknown Number")
		}
	}
}
