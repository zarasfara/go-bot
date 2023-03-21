package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		bot: bot,
	}
}

func (b Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil { // If we got a message

			if update.Message.IsCommand() {
				b.HandleCommand(update.Message)
				continue
			}

			b.HandleMessage(update.Message)
		}
	}
}
