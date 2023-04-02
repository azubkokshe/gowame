package main

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// замените YOUR_TOKEN на токен вашего бота
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Отправь мне номер WhatsApp.")
			bot.Send(msg)
		default:
			phone := strings.ReplaceAll(update.Message.Text, "+", "")
			phone = strings.ReplaceAll(phone, "-", "")
			phone = strings.ReplaceAll(phone, "(", "")
			phone = strings.ReplaceAll(phone, ")", "")
			phone = strings.ReplaceAll(phone, " ", "")
			if _, err := strconv.Atoi(phone); err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректный номер телефона.")
				bot.Send(msg)
			} else {
				waLink := "https://wa.me/" + phone
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, waLink)
				bot.Send(msg)
			}
		}
	}
}
