package main

import (
	"fmt"
	"time"
)

/*
Конструкция c <- "ping" означает отправку "ping",
а msg := <- c — его получение и сохранение в переменную msg.

Данное использование каналов позволяет синхронизировать две горутины.
Когда pinger пытается послать сообщение в канал, он ожидает,
пока printer будет готов получить сообщение. Такое поведение называется блокирующим.
*/
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	c := make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
