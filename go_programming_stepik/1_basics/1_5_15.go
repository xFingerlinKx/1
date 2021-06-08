/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

Входные данные
На вход программе подается целое число d (0 < d < 360).

Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.
*/
package main

import "fmt"

func main() {
	var inputInt int
	fmt.Scan(&inputInt)
	hours := inputInt / 30
	minutes := int((float64(inputInt) * 2) - (float64(hours) * 60))
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
