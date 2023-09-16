package responder

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
)

func (b *Bot) handleResponse(response entity.TaskResponse) error {
	if response.Error != "" {
		return errors.New(response.Error)
	}

	msg := tgbotapi.NewMessage(response.ChatID, response.FileLocation)
	b.bot.Send(msg)

	return nil
}
