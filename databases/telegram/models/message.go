package models

import (
	"github.com/higordiego/rabbitmq-palestra/databases/telegram/database"
)

// Telegram struct handler
type Telegram struct {
	IDTelegram int64
	User       string
	Message    string
}

// SaveMessageTelegram - handler struct
func (telegram *Telegram) SaveMessageTelegram() error {
	db, _ := database.Connection()

	insForm, err := db.Prepare("INSERT INTO telegram(user, message, IDTelegram, created_at) VALUES(?,?, ?, now())")

	if err != nil {
		return err
	}

	insForm.Exec(telegram.User, telegram.Message, telegram.IDTelegram)
	return nil
}
