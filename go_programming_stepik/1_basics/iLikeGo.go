// Напишите программу, которая выведет "I like Go!" 3 раза.

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I like Go!\n"
	fmt.Println(strings.Repeat(s, 3))
}
