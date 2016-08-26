package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://unitoss:turtlerate@unitrating2:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"req_udr_task", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body, err := bodyFrom(os.Args)
	if err != nil {
		fmt.Printf("%v", err)
	}

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
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) (string, error) {
	var s string
	if (len(args) < 3) || os.Args[1] == "" {
		return "", errors.New("not enough args")
	} else if os.Args[1] == "N" {
		s = fmt.Sprintf("{\"errortype\":0, \"count\":%s}", os.Args[2])
	} else if os.Args[1] == "T" {
		s = fmt.Sprintf("{\"errortype\":1, \"count\":%s}", os.Args[2])
	} else if os.Args[1] == "U" {
		s = fmt.Sprintf("{\"errortype\":2, \"count\":%s}", os.Args[2])
	} else if os.Args[1] == "F" {
		s = fmt.Sprintf("{\"errortype\":3, \"count\":%s}", os.Args[2])
	}

	return s, nil
}
