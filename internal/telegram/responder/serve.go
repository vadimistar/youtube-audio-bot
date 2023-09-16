package responder

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"io"
	"net/http"
)

func (b *Bot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response, err := b.decodeResponse(r.Body)
	if err != nil {
		respondError(w, err)
		return
	}

	if response.Error != "" {
		msg := tgbotapi.NewMessage(response.ChatID, "Произошла неизвестная ошибка. Попробуйте позже.")
		b.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(response.ChatID, response.FileLocation)
	b.bot.Send(msg)
}

func (b *Bot) decodeResponse(resp io.Reader) (entity.TaskResponse, error) {
	var decodedResponse entity.TaskResponse

	err := json.NewDecoder(resp).Decode(&decodedResponse)
	if err != nil {
		return entity.TaskResponse{}, ErrInvalidRequestBody
	}

	return decodedResponse, nil
}
