package telegram

import (
	"encoding/json"
	"fmt"
	_ "io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CatImage struct {
	URL string `json:"url"`
}

// команды
const (
	commandStart      = "start"
	commandGetPicture = "getpicture"
)

// обработка команды
func (b Bot) HandleCommand(message *tgbotapi.Message) error {
	var msg tgbotapi.Chattable

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

		fileUrl := tgbotapi.FileURL(body)
		msg = tgbotapi.NewPhoto(message.Chat.ID, fileUrl)

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
func (b Bot) getImage() (string, error) {
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

	var catImages []CatImage
	err = json.NewDecoder(resp.Body).Decode(&catImages)
	if err != nil {
		log.Panic(err)
	}

	return catImages[0].URL, err
}
