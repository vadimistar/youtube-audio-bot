package worker

import (
	"context"
	"github.com/vadimistar/youtube-audio-bot/internal/repository"
	"log/slog"
	"os"
)

type downloadAudio func(ctx context.Context, videoID string) (*os.File, error)

type Worker struct {
	repo          repository.Repository
	downloadAudio downloadAudio
	botURL        string
	logger        *slog.Logger
}

func New(repo repository.Repository, downloadAudio downloadAudio, botURL string, logger *slog.Logger) *Worker {
	return &Worker{repo: repo, downloadAudio: downloadAudio, botURL: botURL, logger: logger}
}
