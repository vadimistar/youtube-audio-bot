package webhook

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vadimistar/youtube-audio-bot/internal/queue"
)

type Bot struct {
	bot   *tgbotapi.BotAPI
	queue queue.Queue
}

func NewBot(bot *tgbotapi.BotAPI, queue queue.Queue) *Bot {
	return &Bot{bot: bot, queue: queue}
}
