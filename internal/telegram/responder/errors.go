package responder

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
)

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
)

func (b *Bot) handleError(w http.ResponseWriter, chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, msgUnknownError)

	switch {
	case errors.Is(err, ErrInvalidRequestBody):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	b.bot.Send(msg)
}
