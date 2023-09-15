package webhook

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot       *tgbotapi.BotAPI
	workerURL string
}

func NewBot(bot *tgbotapi.BotAPI, workerURL string) *Bot {
	return &Bot{bot: bot, workerURL: workerURL}
}
