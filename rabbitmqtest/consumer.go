package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	// Затем мы создаем канал Go, который мы назовем forever, который будет поддерживать
	// работу нашего приложения-потребителя неограниченное время, чтобы оно продолжало
	// потреблять новые сообщения по мере их публикации в очереди.
	forever := make(chan bool)
	// Ниже мы создаем goroutine, которая обрабатывает сообщения, получаемые из нашей TestQueue.
	// В данном случае мы просто выводим содержимое сообщения.
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	// вызываем <-forever, который блокирует завершение нашей основной функции до тех пор,
	// пока канал не будет удовлетворен
	<-forever
}
