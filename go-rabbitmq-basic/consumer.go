package main

import (
	"github.com/streadway/amqp"
	logger "go-rabbitmq/util"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	logger.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	logger.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// create a message receiving queue
	q, err := ch.QueueDeclare(
		"test-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	logger.FailOnError(err, "Failed to declare a queue")

	// build the consumer
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	logger.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever
}
