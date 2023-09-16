package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/vadimistar/youtube-audio-bot/internal/queue/messagequeue"
	"github.com/vadimistar/youtube-audio-bot/internal/repository/objectstorage"
	"github.com/vadimistar/youtube-audio-bot/internal/worker"
	"github.com/vadimistar/youtube-audio-bot/internal/ytdlp"
	"log"
	"net/url"
	"os"
)

func main() {
	godotenv.Load()

	repository, err := objectstorage.NewClient(context.Background(), os.Getenv("YC_BUCKET_NAME"))
	if err != nil {
		log.Fatalf("create repository: %s", err)
	}

	queue, err := messagequeue.NewClient(context.Background(), os.Getenv("MESSAGE_QUEUE_URL"))
	if err != nil {
		log.Fatalf("create message queue: %s", err)
	}

	telegramBotURL, err := url.JoinPath(os.Getenv("TELEGRAM_BOT_CONTAINER_URL"), "responder")
	if err != nil {
		log.Fatalf("join path: %s", err)
	}

	w := worker.New(repository, queue, ytdlp.DownloadAudio, telegramBotURL)

	err = w.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
