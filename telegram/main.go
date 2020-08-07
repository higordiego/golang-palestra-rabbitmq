package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/higordiego/rabbitmq-palestra/telegram/models"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1322251016:AAGgRrIyWBoegNrEjtN-kftnt5QFYu-z_CA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	log.Println("Start application gateway telegram")

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if len(update.Message.Text) > 0 {
			var telegram models.Telegram
			telegram.IDTelegram = update.Message.Chat.ID
			telegram.User = update.Message.From.UserName
			telegram.Message = update.Message.Text

			telegram.RabbitMQQueue()
		}
	}
}
