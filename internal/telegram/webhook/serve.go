package webhook

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func (b *Bot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	update, err := b.decodeUpdate(r)
	if err != nil {
		log.Printf("webhook handler: %s", err)
		return
	}

	b.handleUpdate(update)
}

func (b *Bot) decodeUpdate(r *http.Request) (update tgbotapi.Update, err error) {
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
			err = errors.Wrap(closeErr, "decode update")
		}
	}()

	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		return update, errors.Wrap(err, "decode update")
	}

	return update, nil
}
