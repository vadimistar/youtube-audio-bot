package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/vadimistar/youtube-audio-bot/internal/telegram"
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

	telegramBot := telegram.NewBot(bot)
	mux := http.NewServeMux()
	mux.HandleFunc("/", telegramBot.WebhookHandler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
