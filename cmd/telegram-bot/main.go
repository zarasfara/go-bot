package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/zarasfara/telegram-bot-api/internal/telegram"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_KEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	telegramBot := telegram.NewBot(bot)

	telegramBot.Start()

}
