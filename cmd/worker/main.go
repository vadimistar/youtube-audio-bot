package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"github.com/vadimistar/youtube-audio-bot/internal/repository/objectstorage"
	"github.com/vadimistar/youtube-audio-bot/internal/worker"
	"github.com/vadimistar/youtube-audio-bot/internal/ytdlp"
	"log"
	"log/slog"
	"net/url"
	"os"
)

func main() {
	godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	repository, err := objectstorage.NewClient(context.Background(), os.Getenv("YC_BUCKET_NAME"))
	if err != nil {
		log.Fatalf("create repository: %s", err)
	}

	telegramBotURL, err := url.JoinPath(os.Getenv("TELEGRAM_BOT_CONTAINER_URL"), "responder")
	if err != nil {
		log.Fatalf("join path: %s", err)
	}

	w := worker.New(repository, ytdlp.DownloadAudio, telegramBotURL, logger)

	e := echo.New()

	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	e.POST("/", w.HTTP)

	log.Fatalln(e.Start(":" + os.Getenv("PORT")))
}
