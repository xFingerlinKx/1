package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func failOnError1(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 1) connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError1(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 2) create a channel
	ch, err := conn.Channel()
	failOnError1(err, "Failed to open a channel")
	defer ch.Close()

	// 3) to send, we must declare a queue for us to send to
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError1(err, "Failed to declare a queue")

	body := bodyFrom(os.Args)
	// 4) publish a message to the queue
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError1(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
