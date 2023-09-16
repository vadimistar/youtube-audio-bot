package worker

import (
	"context"
	"github.com/vadimistar/youtube-audio-bot/internal/queue"
	"github.com/vadimistar/youtube-audio-bot/internal/repository"
	"os"
)

type downloadAudio func(ctx context.Context, videoID string) (*os.File, error)

type Worker struct {
	repo          repository.Repository
	queue         queue.Queue
	downloadAudio downloadAudio
	botURL        string
}

func New(repo repository.Repository, queue queue.Queue, downloadAudio downloadAudio, botURL string) *Worker {
	return &Worker{repo: repo, queue: queue, downloadAudio: downloadAudio, botURL: botURL}
}
