package models

import (
	"encoding/json"
	"log"

	"github.com/higordiego/rabbitmq-palestra/business/whatsapp/database"
	rabbit "github.com/higordiego/rabbitmq-palestra/business/whatsapp/integration"
)

// Whatsapp struct handler
type Whatsapp struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

// SendChannelComunicate - handler send channel telegram
func (telegram *Whatsapp) SendChannelComunicate() error {
	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")

	defer conn.Conn.Close()

	if err != nil {
		return err
	}

	jsonString, _ := json.Marshal(&telegram)
	conn.PublishRabbitMQ("messages-whatsapp-comunicate", []byte(jsonString), "whatsapp")
	return nil
}

// SendErrorChannelWhatsapp - handler send change telegram
func (telegram *Whatsapp) SendErrorChannelWhatsapp() error {
	conn, err := rabbit.GetConnectionRabbit("amqp://guest:guest@localhost")

	defer conn.Conn.Close()

	if err != nil {
		return err
	}

	jsonString, _ := json.Marshal(&telegram)

	log.Println(telegram)

	conn.PublishRabbitMQ("business-error-whatsapp", []byte(jsonString), "whatsapp")
	return nil
}

// SaveMessageWhatsapp - handler struct
func (whatsapp *Whatsapp) SaveMessageWhatsapp() error {

	whatsapp.SendChannelComunicate()

	db, _ := database.Connection()

	insForm, err := db.Prepare("INSERT INTO whatsapp(name, phone, message, created_at) VALUES(?,?,?, now())")

	if err != nil {
		log.Println("Erro em se conectar ao banco de dados")
		whatsapp.SendErrorChannelWhatsapp()
		return err
	}

	insForm.Exec(whatsapp.Name, whatsapp.Phone, whatsapp.Message)
	return nil
}
