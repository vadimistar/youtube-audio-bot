package responder

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vadimistar/youtube-audio-bot/internal/repository"
	"log/slog"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	repo        repository.Repository
	logger      *slog.Logger
	ycBucketURL string
}

func NewBot(bot *tgbotapi.BotAPI, repo repository.Repository, logger *slog.Logger, ycBucketURL string) *Bot {
	return &Bot{bot: bot, repo: repo, logger: logger, ycBucketURL: ycBucketURL}
}
