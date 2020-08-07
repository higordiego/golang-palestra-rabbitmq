package models

import (
	"github.com/higordiego/rabbitmq-palestra/databases/whatsapp/database"
)

// Whatsapp struct handler
type Whatsapp struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

// SaveMessageWhatsapp - handler struct
func (whatsapp *Whatsapp) SaveMessageWhatsapp() error {

	db, _ := database.Connection()

	insForm, err := db.Prepare("INSERT INTO whatsapp(name, phone, message) VALUES(?,?, ?)")

	if err != nil {
		return err
	}

	insForm.Exec(whatsapp.Name, whatsapp.Phone, whatsapp.Message)
	return nil
}
