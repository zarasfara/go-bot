package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errMessage = errors.New("ошибка отправки сообщения")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, err.Error())
	switch err {
	case errMessage:
		msg.Text = "ошибка отправки сообщения"
	default:
		msg.Text = "неизвестная ошибка"
	}
	b.bot.Send(msg)
}
