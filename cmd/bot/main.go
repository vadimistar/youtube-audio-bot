package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/vadimistar/youtube-audio-bot/internal/queue/messagequeue"
	"github.com/vadimistar/youtube-audio-bot/internal/repository/objectstorage"
	"github.com/vadimistar/youtube-audio-bot/internal/telegram/responder"
	"github.com/vadimistar/youtube-audio-bot/internal/telegram/webhook"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

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

	webhookBot := webhook.NewBot(bot, queue)
	responderBot := responder.NewBot(bot, repository)

	mux := http.NewServeMux()
	mux.Handle("/", webhookBot)
	mux.Handle("/responder", responderBot)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
