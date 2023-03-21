package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
)

const (
	commandStart      = "start"
	commandGetPicture = "getpicture"
)

func (b Bot) HandleCommand(message *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig

	// сообщение - команда
	switch message.Command() {
	case commandStart:
		msg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Привет, %s", message.Chat.FirstName))
	case commandGetPicture:
		body, err := b.HandleImage()
		if err != nil {
			log.Fatalf("[%s] %s", message.From.UserName, err)
		}
		msg = tgbotapi.NewMessage(message.Chat.ID, string(body))

	default:
		msg = tgbotapi.NewMessage(message.Chat.ID, "Я такой команды не знаю :[")
	}

	_, err := b.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func (b Bot) HandleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}

func (b Bot) HandleImage() ([]byte, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body, err

}
