package webhook

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	ErrInvalidURL = errors.New("invalid url")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Неизвестная ошибка. Попробуйте позже.")

	switch {
	case errors.Is(err, ErrInvalidURL):
		msg.Text = "Неправильная ссылка."
	default:
		log.Printf("handleError: chatID=%d err=%s", chatID, err)
	}

	_, err = b.bot.Send(msg)
	if err != nil {
		log.Printf("handleError: send error message: %s", err)
	}
}
