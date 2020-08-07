package models

import (
	"encoding/json"

	rabbit "github.com/higordiego/rabbitmq-palestra/telegram/integration"
)

// Telegram struct handler
type Telegram struct {
	IDTelegram      int64
	User    string
	Message string
}

// RabbitMQQueue - handler message RabbitMQ
func (telegram *Telegram) RabbitMQQueue() error {

	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")

	defer conn.Conn.Close()

	if err != nil {
		return err
	}

	jsonString, _ := json.Marshal(&telegram)

	conn.PublishRabbitMQ("messages-telegram", []byte(jsonString), "telegram")
	return nil
}
