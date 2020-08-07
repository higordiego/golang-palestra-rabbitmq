package main

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"

	rabbit "github.com/higordiego/rabbitmq-palestra/consumers/business/integration"
	"github.com/higordiego/rabbitmq-palestra/consumers/business/models"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")
	if err != nil {
		panic(err)
	}

	err = conn.StartQueeConsumer("text-telegram", "messages-telegram", handler, 1)

	if err != nil {
		panic(err)
	}

	log.Println("Consumer bussiness telegram application start")

	forever := make(chan bool)
	<-forever
}

func handler(d amqp.Delivery) bool {
	if d.Body == nil {
		log.Println("Error, no message body!")
		return false
	}

	log.Println("business telegram recebe", string(d.Body))
	var telegram models.Telegram

	err := json.Unmarshal(d.Body, &telegram)
	if err != nil {
		panic(err)
	}

	telegram.SaveMessageTelegram()
	return true
}
