package responder

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"log/slog"
	"net/http"
	url2 "net/url"
)

const (
	msgUnknownError = "Произошла неизвестная ошибка. Попробуйте позже."
)

func (b *Bot) HTTP(c echo.Context) error {
	resp := new(entity.TaskResponse)
	err := c.Bind(resp)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid input")
	}

	logger := slog.With(slog.Int64("chatID", resp.ChatID))

	if resp.Error != "" {
		logger.Error("response error", slog.String("error", resp.Error))
		msg := tgbotapi.NewMessage(resp.ChatID, msgUnknownError)
		b.bot.Send(msg)
		return c.String(http.StatusInternalServerError, resp.Error)
	}

	audio, err := b.repo.Get(c.Request().Context(), resp.Key)
	if err != nil {
		logger.Error("response error", slog.String("error", resp.Error))
		return c.String(http.StatusBadRequest, "invalid key")
	}

	defer audio.Close()

	url, err := url2.JoinPath(b.ycBucketURL, resp.Key)
	if err != nil {
		logger.Error("join path for yc bucket url", slog.String("error", resp.Error))
		return c.String(http.StatusBadRequest, "invalid url")
	}
	msg := tgbotapi.NewMessage(resp.ChatID, "Итоговый файл (он будет удален через 24 часа)\n"+url)
	_, err = b.bot.Send(msg)
	if err != nil {
		logger.Error("send response to user", slog.String("error", err.Error()))
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return nil
}
