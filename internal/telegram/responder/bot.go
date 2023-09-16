package responder

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vadimistar/youtube-audio-bot/internal/repository"
)

type Bot struct {
	bot  *tgbotapi.BotAPI
	repo repository.Repository
}

func NewBot(bot *tgbotapi.BotAPI, repo repository.Repository) *Bot {
	return &Bot{bot: bot, repo: repo}
}
