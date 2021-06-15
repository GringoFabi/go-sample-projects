package main

import (
	"github.com/streadway/amqp"
	"go-rabbitmq-advanced/util"
	"log"
	"os"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")

	body := util.BodyFrom(os.Args)
	err = ch.Publish(
		"",           // exchange
		q.Name,       // routing key
		false,        // mandatory
		false,
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	util.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
