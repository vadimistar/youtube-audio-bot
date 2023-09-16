package webhook

import (
	"bytes"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"io"
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
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваше видео обрабатывается. Как только оно будет готово, я отправлю вам итоговый файл.")
	b.bot.Send(msg)
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

	err = b.sendTaskToWorker(entity.TaskRequest{ChatID: message.Chat.ID, VideoURL: inputURL})
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) sendTaskToWorker(task entity.TaskRequest) (err error) {
	defer func() {
		err = errors.Wrap(err, "send task to worker")
	}()

	serializedTask, err := serializeTask(task)
	if err != nil {
		return err
	}

	_, err = http.Post(b.workerURL, "application/json", serializedTask)
	if err != nil {
		return err
	}

	return nil
}

func serializeTask(task entity.TaskRequest) (io.Reader, error) {
	serializedTask := new(bytes.Buffer)

	err := json.NewEncoder(serializedTask).Encode(task)
	if err != nil {
		return nil, err
	}

	return serializedTask, nil
}

func validateURL(text string) error {
	_, err := url.ParseRequestURI(text)
	if err != nil {
		return ErrInvalidURL
	}
	return nil
}
