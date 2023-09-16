package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"log"
)

func (w *Worker) Start() error {
	messages, err := w.queue.Receive(context.Background())
	if err != nil {
		return err
	}
	log.Printf("received %d messages", len(messages))

	if len(messages) == 0 {
		return nil
	}

	for _, message := range messages {
		request, err := deserialize(message)
		if err != nil {
			return err
		}

		err = w.processRequest(request)
		if err != nil {
			return err
		}
	}

	return nil
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
