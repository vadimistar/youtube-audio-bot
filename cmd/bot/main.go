package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"github.com/vadimistar/youtube-audio-bot/internal/queue/messagequeue"
	"github.com/vadimistar/youtube-audio-bot/internal/repository/objectstorage"
	"github.com/vadimistar/youtube-audio-bot/internal/telegram/responder"
	"github.com/vadimistar/youtube-audio-bot/internal/telegram/webhook"
	"log"
	"log/slog"
	"os"
)

func main() {
	godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatalf("new telegram bot API: %s", err)
	}

	repository, err := objectstorage.NewClient(context.Background(), os.Getenv("YC_BUCKET_NAME"))
	if err != nil {
		log.Fatalf("create repository: %s", err)
	}

	queue, err := messagequeue.NewClient(context.Background(), os.Getenv("MESSAGE_QUEUE_URL"))
	if err != nil {
		log.Fatalf("create message queue: %s", err)
	}

	webhookBot := webhook.NewBot(bot, queue, logger)
	responderBot := responder.NewBot(bot, repository, logger, os.Getenv("YC_BUCKET_URL"))

	e := echo.New()

	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	e.POST("/", webhookBot.HTTP)
	e.POST("/responder", responderBot.HTTP)

	log.Fatalln(e.Start(":" + os.Getenv("PORT")))
}
