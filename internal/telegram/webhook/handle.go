package webhook

import (
	"bytes"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"net/http"
	"net/url"
	"strings"
)

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	err := b.handleMessage(update.Message)
	if err != nil {
		b.handleError(update.Message.Chat.ID, err)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) (err error) {
	defer func() {
		err = errors.Wrap(err, "handle message")
	}()

	inputURL := strings.TrimSpace(message.Text)

	err = validateURL(inputURL)
	if err != nil {
		return err
	}

	err = b.sendTaskToWorker(message.Chat.ID, inputURL)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) sendTaskToWorker(chatID int64, url string) (err error) {
	defer func() {
		err = errors.Wrap(err, "send task to worker")
	}()

	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(entity.TaskRequest{
		ChatID:   chatID,
		VideoURL: url,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, b.workerURL, reqBody)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func validateURL(text string) error {
	_, err := url.Parse(text)
	if err != nil {
		return ErrInvalidURL
	}
	return nil
}
