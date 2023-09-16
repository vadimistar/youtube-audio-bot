package worker

import (
	"bytes"
	"encoding/json"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"log"
	"net/http"
)

func (w *Worker) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	messages, err := receiveMessages(r.Body)
	if err != nil {
		log.Printf("receive messages: %s", err)
	}
	log.Printf("received %d messages", len(messages.Messages))

	err = w.processMessages(messages)
	if err != nil {
		log.Printf("process messages: %s", err)
	}
}

func deserialize(msg string) (request entity.TaskRequest, err error) {
	var msgBuffer bytes.Buffer
	msgBuffer.WriteString(msg)

	err = json.NewDecoder(&msgBuffer).Decode(&request)
	if err != nil {
		return entity.TaskRequest{}, err
	}
	return request, nil
}
