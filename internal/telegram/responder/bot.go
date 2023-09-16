package responder

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
)

type storage interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Put(ctx context.Context, key string, file io.Reader) error
}

type Bot struct {
	bot     *tgbotapi.BotAPI
	storage storage
}

func NewBot(bot *tgbotapi.BotAPI, storage storage) *Bot {
	return &Bot{bot: bot, storage: storage}
}
