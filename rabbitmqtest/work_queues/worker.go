package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failError1(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 1) connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failError1(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	// 2) create a channel
	ch, err := conn.Channel()
	failError1(err, "Failed to open a channel")
	defer ch.Close()

	/*
		Note that we declare the queue here, as well. Because we might start the consumer before
		the publisher, we want to make sure the queue exists before we try to consume messages from it.
		We're about to tell the server to deliver us the messages from the queue.
		Since it will push us messages asynchronously, we will read the messages
		from a channel (returned by amqp::Consume) in a goroutine.
	*/
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failError1(err, "Failed to declare a queue")
	/*
		Вы могли заметить, что отправка по-прежнему работает не так, как нам хотелось бы.
		Например, в ситуации с двумя воркерами, когда все нечетные сообщения тяжелые,
		а четные - легкие, один воркер будет постоянно занят, а другой почти не будет
		выполнять никакой работы. Что ж, RabbitMQ ничего об этом не знает и по-прежнему
		будет отправлять сообщения равномерно.

		Это происходит потому, что RabbitMQ просто отправляет сообщение, когда оно попадает в очередь.
		Он не смотрит на количество неподтвержденных сообщений для потребителя. Он просто слепо
		отправляет каждое n-е сообщение n-му потребителю.

		Чтобы избежать этого, мы можем установить для счетчика предварительной выборки значение 1.
		Это говорит RabbitMQ не отправлять более одного сообщения рабочему за раз. Или, другими словами,
		не отправляйте новое сообщение работнику, пока он не обработает и не подтвердит предыдущее.
		Вместо этого он отправит его следующему работнику, который еще не занят.
	*/
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failError1(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failError1(err, "Failed to register a consumer")
	log.Printf("q.Name: %s", q.Name)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
