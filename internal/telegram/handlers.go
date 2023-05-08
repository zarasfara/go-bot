package telegram

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// команды
const (
	commandStart      = "start"
	commandGetPicture = "getpicture"
)

// обработка команды
func (b Bot) HandleCommand(message *tgbotapi.Message) error {
	var msg tgbotapi.MessageConfig

	// сообщение - команда
	switch message.Command() {
	case commandStart:
		msg = tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Привет, %s", message.Chat.FirstName))
	case commandGetPicture:
		body, err := b.getImage()
		if err != nil {
			return err
			// log.Fatalf("[%s] %s", message.From.UserName, err)
		}
		msg = tgbotapi.NewMessage(message.Chat.ID, string(body))

	default:
		msg = tgbotapi.NewMessage(message.Chat.ID, "Я такой команды не знаю :[")
	}

	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

// обработка сообщения
func (b Bot) HandleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	_, err := b.bot.Send(msg)
	if err != nil {
		return errMessage
	}

	return nil
}

// получение изображений
func (b Bot) getImage() ([]byte, error) {

	url := "https://api.thecatapi.com/v1/images/search"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("x-api-key", os.Getenv("CATS_API_KEY"))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	return body, err

}


