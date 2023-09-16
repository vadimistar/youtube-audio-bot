package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/vadimistar/youtube-audio-bot/internal/entity"
	"net/http"
)

func (w *Worker) processRequest(req entity.TaskRequest) error {
	videoID, err := parseYoutubeVideoID(req.VideoURL)
	if err != nil {
		return err
	}

	audio, err := w.downloadAudio(context.Background(), videoID)
	if err != nil {
		return err
	}

	err = w.repo.Put(context.Background(), audio.Name(), audio)
	if err != nil {
		return err
	}

	err = w.sendResponse(req.ChatID, audio.Name())
	if err != nil {
		return err
	}

	return nil
}

func (w *Worker) sendResponse(chatID int64, key string) error {
	var responseBuffer bytes.Buffer
	err := json.NewEncoder(&responseBuffer).Encode(entity.TaskResponse{
		ChatID: chatID,
		Key:    key,
	})
	if err != nil {
		return err
	}

	_, err = http.Post(w.botURL, "application/json", &responseBuffer)
	if err != nil {
		return err
	}

	return nil
}
