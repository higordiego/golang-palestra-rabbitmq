package main

import (
	"fmt"
	"log"

	rabbit "github.com/higordiego/rabbitmq-palestra/consumers/whatsapp/integration"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")
	if err != nil {
		panic(err)
	}

	err = conn.StartQueeConsumer("text-whatsapp-comunicate", "messages-whatsapp-comunicate", handler, 100)

	if err != nil {
		panic(err)
	}

	fmt.Println("Consumer whatsapp application start")

	forever := make(chan bool)
	<-forever
}

func handler(d amqp.Delivery) bool {
	if d.Body == nil {
		log.Println("Error, no message body!")
		return false
	}
	log.Println(string(d.Body))
	return true
}
