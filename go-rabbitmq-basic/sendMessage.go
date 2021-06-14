package main

import (
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	logger "go-rabbitmq/util"
	"log"
	"os"
)

func main() {
	// receive message from terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your message: ")
	mPayload, _ := reader.ReadString('\n')

	// connect to rabbitmq
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

	// publishing a new message
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(mPayload),
		},
	)
	logger.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent message: %s", mPayload)
}
