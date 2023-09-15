package telegram

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var update tgbotapi.Update

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Printf("invalid request body (expect tgbotapi.Message): %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	_, err = b.bot.Send(msg)
	if err != nil {
		log.Printf("cannot send message to %s: %s", update.Message.Chat.ID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
