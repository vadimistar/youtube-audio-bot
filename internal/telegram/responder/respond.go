package responder

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"net/http"
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
