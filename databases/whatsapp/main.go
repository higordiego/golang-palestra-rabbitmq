package main

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"

	rabbit "github.com/higordiego/rabbitmq-palestra/databases/whatsapp/integration"
	"github.com/higordiego/rabbitmq-palestra/databases/whatsapp/models"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")
	if err != nil {
		panic(err)
	}

	err = conn.StartQueeConsumer("text-whatsapp-error-comunicate", "business-error-whatsapp", handler, 1)

	if err != nil {
		panic(err)
	}

	log.Println("Consumer database error whatsapp application start")

	forever := make(chan bool)
	<-forever
}

func handler(d amqp.Delivery) bool {
	if d.Body == nil {
		log.Println("Error, no message body!")
		return false
	}
	var whatsapp models.Whatsapp

	json.Unmarshal(d.Body, &whatsapp)

	err := whatsapp.SaveMessageWhatsapp()
	if err != nil {
		return false
	}

	return true
}
