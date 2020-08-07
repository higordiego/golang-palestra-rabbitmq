package main

import (
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"

	rabbit "github.com/higordiego/rabbitmq-palestra/databases/telegram/integration"
	"github.com/higordiego/rabbitmq-palestra/databases/telegram/models"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")
	if err != nil {
		panic(err)
	}

	err = conn.StartQueeConsumer("text-telegram-error-comunicate", "business-error-telegram", handler, 1)

	if err != nil {
		panic(err)
	}

	log.Println("Consumer database error telegram application start")

	forever := make(chan bool)
	<-forever
}

func handler(d amqp.Delivery) bool {
	if d.Body == nil {
		log.Println("Error, no message body!")
		return false
	}
	var telegram models.Telegram

	json.Unmarshal(d.Body, &telegram)

	err := telegram.SaveMessageTelegram()
	if err != nil {
		return false
	}

	return true
}
