package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	commandStart = "start"
)

var name = map[string]string{
	"random_image": "Случайное изображение",
}

var KeyBoardCommand = []string{"Случайное изображение"}

func (b Bot) HandleCommand(message *tgbotapi.Message) {

	//сообщение - команда
	switch message.Command() {
	case commandStart:
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Привет, %s", message.Chat.FirstName))
		//открываем клавиатуру
		msg.ReplyMarkup = numericKeyboard

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

func (b Bot) HandleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	if containString(KeyBoardCommand, message.Text) {
		b.handleKeyBoard(message)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		_, err := b.bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (b Bot) handleKeyBoard(message *tgbotapi.Message) {

	switch message.Text {
	case name["random_image"]:
		b.handleImageKey(message)
	}

}

func (b Bot) handleImageKey(message *tgbotapi.Message) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, string(body))
	_, err = b.bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}

}
