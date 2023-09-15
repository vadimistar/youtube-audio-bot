package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
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

	telegramBot := webhook.NewBot(bot, os.Getenv("WORKER_CONTAINER_URL"))

	mux := http.NewServeMux()
	mux.Handle("/", telegramBot)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
