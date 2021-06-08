package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 1) connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 2) create a channel
	ch, err := conn.Channel()
	failError(err, "Failed to open a channel")
	defer ch.Close()

	/*
		Note that we declare the queue here, as well. Because we might start the consumer before
		the publisher, we want to make sure the queue exists before we try to consume messages from it.
		We're about to tell the server to deliver us the messages from the queue.
		Since it will push us messages asynchronously, we will read the messages
		from a channel (returned by amqp::Consume) in a goroutine.
	*/
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"", // consumer
		true, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	failError(err, "Failed to register a consumer")
	// Затем мы создаем канал Go, который мы назовем forever, который будет поддерживать
	// работу нашего приложения-потребителя неограниченное время, чтобы оно продолжало
	// потреблять новые сообщения по мере их публикации в очереди.
	forever := make(chan bool)
	// Ниже мы создаем goroutine, которая обрабатывает сообщения, получаемые из нашей TestQueue.
	// В данном случае мы просто выводим содержимое сообщения.
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// вызываем <-forever, который блокирует завершение нашей основной функции до тех пор,
	// пока канал не будет удовлетворен
	<-forever
}
