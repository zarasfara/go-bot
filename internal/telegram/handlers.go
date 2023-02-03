package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const (
	commandStart = "start"
)

func (b Bot) HandleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	_, err := b.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}

}

func (b Bot) HandleCommand(message *tgbotapi.Message) {

	switch message.Command() {
	case commandStart:
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Привет, %s", message.Chat.FirstName))
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Я такой команды не знаю :[")
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

}
