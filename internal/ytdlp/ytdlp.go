package ytdlp

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func DownloadAudio(ctx context.Context, videoID string) (*os.File, error) {
	cmd := exec.CommandContext(ctx, "yt-dlp", "--extract-audio", "--audio-format", "mp3", "-o", "%(id)s.%(ext)s", videoID)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("execute yt-dlp error = %s; output = %s", err, output)
	}

	return os.Open(videoID + ".mp3")
}
