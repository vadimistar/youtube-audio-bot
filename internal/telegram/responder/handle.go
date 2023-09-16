package responder

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
)

func (b *Bot) handleResponse(response entity.TaskResponse) error {
	if response.Error != "" {
		return errors.New(response.Error)
	}

	audio, err := b.repo.Get(context.Background(), response.Key)
	if err != nil {
		return err
	}

	defer audio.Close()

	msg := tgbotapi.NewAudio(response.ChatID, tgbotapi.FileReader{
		Name:   response.Key,
		Reader: audio,
	})
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
