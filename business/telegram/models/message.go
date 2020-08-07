package models

import (
	"encoding/json"
	"log"

	"github.com/higordiego/rabbitmq-palestra/consumers/business/database"
	rabbit "github.com/higordiego/rabbitmq-palestra/consumers/business/integration"
)

// Telegram struct handler
type Telegram struct {
	IDTelegram int64
	User       string
	Message    string
}

// SendChannelComunicate - handler send channel telegram
func (telegram *Telegram) SendChannelComunicate() error {
	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")

	defer conn.Conn.Close()

	if err != nil {
		return err
	}

	jsonString, _ := json.Marshal(&telegram)
	conn.PublishRabbitMQ("messages-telegram-comunicate", []byte(jsonString), "telegram")
	return nil
}

// SendErrorChannelTelegram - handler send change telegram
func (telegram *Telegram) SendErrorChannelTelegram() error {
	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")

	defer conn.Conn.Close()

	if err != nil {
		return err
	}

	jsonString, _ := json.Marshal(&telegram)

	conn.PublishRabbitMQ("business-error-telegram", []byte(jsonString), "telegram")
	return nil
}

// SaveMessageTelegram - handler struct
func (telegram *Telegram) SaveMessageTelegram() error {

	telegram.SendChannelComunicate()

	db, _ := database.Connection()

	insForm, err := db.Prepare("INSERT INTO telegram(user, message, IDTelegram, created_at) VALUES(?,?, ?, now())")

	if err != nil {
		log.Println("Erro na conexão com o banco de dados!")
		telegram.SendErrorChannelTelegram()
		return err
	}

	insForm.Exec(telegram.User, telegram.Message, telegram.IDTelegram)
	return nil
}
