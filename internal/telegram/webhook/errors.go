package webhook

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	ErrInvalidURL = errors.New("invalid url")
)

const (
	MsgUnknownError = "Неизвестная ошибка. Попробуйте позже."
	MsgInvalidURL   = "Неправильная ссылка."
)

func (b *Bot) sendError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, MsgUnknownError)

	switch {
	case errors.Is(err, ErrInvalidURL):
		msg.Text = MsgInvalidURL
	default:
		log.Printf("sendError: chatID=%d err=%s", chatID, err)
	}

	b.bot.Send(msg)
}
