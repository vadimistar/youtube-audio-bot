package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/vadimistar/youtube-audio-bot/internal/objectstorage"
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

	storage, err := objectstorage.NewClient(context.Background(), os.Getenv("YC_BUCKET_NAME"))
	if err != nil {
		log.Fatalf("create object storage client: %s", err)
	}

	webhookBot := webhook.NewBot(bot, os.Getenv("WORKER_CONTAINER_URL"))
	responderBot := responder.NewBot(bot, storage)

	mux := http.NewServeMux()
	mux.Handle("/", webhookBot)
	mux.Handle("/responder", responderBot)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
