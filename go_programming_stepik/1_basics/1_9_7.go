/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого
сумма первых трёх цифр совпадает с суммой трёх последних.

Формат входных данных
На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".

Альтернативное решение:
func main() {
    var s string
    fmt.Scan(&s)
    if s[0] + s[1] + s[2] == s[3] + s[4] + s[5] {
        fmt.Print("YES")
    } else {
        fmt.Print("NO")
    }
}

*/

package main

import "fmt"

func main() {
	var inputInt int
	fmt.Scan(&inputInt)

	number1 := inputInt / 100000
	number2 := (inputInt / 10000) % 10
	number3 := (inputInt / 1000) % 10 % 10
	number4 := (inputInt / 100) % 10 % 10 % 10
	number5 := (inputInt / 10) % 10 % 10 % 10 % 10
	number6 := inputInt % 10 % 10 % 10 % 10
	firstThreeNumSum := number1 + number2 + number3
	secondThreeNumSum := number4 + number5 + number6
	if firstThreeNumSum == secondThreeNumSum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
