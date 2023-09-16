package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/vadimistar/youtube-audio-bot/internal/repository/objectstorage"
	"github.com/vadimistar/youtube-audio-bot/internal/worker"
	"github.com/vadimistar/youtube-audio-bot/internal/ytdlp"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	godotenv.Load()

	repository, err := objectstorage.NewClient(context.Background(), os.Getenv("YC_BUCKET_NAME"))
	if err != nil {
		log.Fatalf("create repository: %s", err)
	}

	telegramBotURL, err := url.JoinPath(os.Getenv("TELEGRAM_BOT_CONTAINER_URL"), "responder")
	if err != nil {
		log.Fatalf("join path: %s", err)
	}

	w := worker.New(repository, ytdlp.DownloadAudio, telegramBotURL)

	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"), w))
}
