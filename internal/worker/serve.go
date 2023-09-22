package worker

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"log/slog"
	"net/http"
)

func (w *Worker) HTTP(c echo.Context) error {
	messages := new(messages)
	err := c.Bind(messages)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid input")
	}
	w.logger.Info("received messages", slog.Int("count", len(messages.Messages)))

	for _, message := range messages.Messages {
		request, err := deserialize(message.Details.Message.Body)
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid message body")
		}

		logger := w.logger.With(slog.Int64("chatID", request.ChatID), slog.String("videoID", request.VideoID))

		audio, err := w.downloadAudio(c.Request().Context(), request.VideoID)
		if err != nil {
			logger.Error("download audio", slog.String("error", err.Error()))
			return c.String(http.StatusInternalServerError, "internal server error")
		}

		err = w.repo.Put(c.Request().Context(), audio.Name(), audio)
		if err != nil {
			logger.Error("save file into repository", slog.String("error", err.Error()))
			return c.String(http.StatusInternalServerError, "internal server error")
		}

		err = w.sendResponse(request.ChatID, audio.Name())
		if err != nil {
			logger.Error("send response", slog.String("error", err.Error()))
			return c.String(http.StatusInternalServerError, "internal server error")
		}
	}

	return nil
}

func (w *Worker) sendResponse(chatID int64, key string) error {
	var responseBuffer bytes.Buffer
	err := json.NewEncoder(&responseBuffer).Encode(entity.TaskResponse{
		ChatID: chatID,
		Key:    key,
	})
	if err != nil {
		return err
	}

	_, err = http.Post(w.botURL, "application/json", &responseBuffer)
	if err != nil {
		return err
	}

	return nil
}

func deserialize(msg string) (request entity.TaskRequest, err error) {
	var msgBuffer bytes.Buffer
	msgBuffer.WriteString(msg)

	err = json.NewDecoder(&msgBuffer).Decode(&request)
	if err != nil {
		return entity.TaskRequest{}, err
	}
	return request, nil
}
