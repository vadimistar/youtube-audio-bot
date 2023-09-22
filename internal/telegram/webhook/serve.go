package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

func (b *Bot) HTTP(c echo.Context) error {
	update := new(tgbotapi.Update)
	err := c.Bind(update)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid input")
	}

	if update.Message == nil {
		return c.String(http.StatusOK, "not a message")
	}

	logger := b.logger.With(slog.Int64("chatID", update.Message.Chat.ID), slog.String("message", update.Message.Text))

	if update.Message.IsCommand() {
		return c.String(http.StatusOK, "is a command")
	}

	inputURL := strings.TrimSpace(update.Message.Text)

	videoID, err := parseYoutubeVideoID(inputURL)
	if err != nil {
		b.sendError(update.Message.Chat.ID, ErrInvalidURL)
		logger.Error("parse youtube video id: ", slog.String("error", err.Error()))
		return c.String(http.StatusOK, "invalid input")
	}

	task := entity.TaskRequest{ChatID: update.Message.Chat.ID, VideoID: videoID}
	err = b.sendTaskToWorker(c.Request().Context(), task)
	if err != nil {
		logger.Error("send task to worker", slog.String("error", err.Error()))
		return c.String(http.StatusOK, "internal error")
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваше видео обрабатывается.")
	_, err = b.bot.Send(msg)
	if err != nil {
		logger.Error("send success message to user", slog.String("error", err.Error()))
		return c.String(http.StatusOK, "internal error")
	}

	return nil
}

func (b *Bot) sendTaskToWorker(ctx context.Context, task entity.TaskRequest) (err error) {
	defer func() {
		err = errors.Wrap(err, "send task to worker")
	}()

	serializedTask, err := serializeTask(task)
	if err != nil {
		return err
	}

	message, err := io.ReadAll(serializedTask)
	if err != nil {
		return err
	}

	err = b.queue.Send(ctx, string(message))
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
