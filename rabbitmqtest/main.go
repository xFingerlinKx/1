package main

import (
	"fmt"

	"github.com/streadway/amqp"
)
// Rabbit - http://localhost:15672/
func main() {
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}
	// открываем канал к инстансу Rabbit через conn, который мы уже установили
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}

	// defer - функция отложенного вызова
	defer ch.Close()

	// QueueDeclare объявляет очередь для хранения сообщений и доставки их потребителям.
	// Объявление создает очередь, если она еще не существует, или гарантирует, что
	// существующая очередь соответствует тем же параметрам.
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: 		 []byte("Hello World"),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Published Message to Queue")
}