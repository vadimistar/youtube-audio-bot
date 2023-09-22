package webhook

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vadimistar/youtube-audio-bot/internal/queue"
	"log/slog"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	queue  queue.Queue
	logger *slog.Logger
}

func NewBot(bot *tgbotapi.BotAPI, queue queue.Queue, logger *slog.Logger) *Bot {
	return &Bot{bot: bot, queue: queue, logger: logger}
}
