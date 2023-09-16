package responder

import (
	"encoding/json"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"io"
	"log"
	"net/http"
)

const (
	msgUnknownError = "Произошла неизвестная ошибка. Попробуйте позже."
)

func (b *Bot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	response, err := b.decodeResponse(r.Body)
	if err != nil {
		log.Printf("err = %s", err)
		return
	}

	err = b.handleResponse(response)
	if err != nil {
		b.handleError(w, response.ChatID, err)
		return
	}
}

func (b *Bot) decodeResponse(resp io.Reader) (entity.TaskResponse, error) {
	var decodedResponse entity.TaskResponse

	err := json.NewDecoder(resp).Decode(&decodedResponse)
	if err != nil {
		return entity.TaskResponse{}, ErrInvalidRequestBody
	}

	return decodedResponse, nil
}
