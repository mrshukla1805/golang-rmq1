package main

import (
	"log"

	connector "github.com/streadway/amqp"
)

func main() {
	conn, err := connector.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMq", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"first", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	body := "Parth Shukla!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		connector.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish message", err)
	}
	log.Printf(" [x] Sent %s", body)
}
