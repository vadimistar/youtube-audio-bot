package webhook

import (
	"fmt"
	"regexp"
)

var regexYoutubeVideoID = regexp.MustCompile(`.*(?:(?:youtu\.be\/|v\/|vi\/|u\/\w\/|embed\/)|(?:(?:watch)?\?v(?:i)?=|\&v(?:i)?=))([^#\&\?]*).*`)

func parseYoutubeVideoID(youtubeURL string) (videoID string, err error) {
	matches := regexYoutubeVideoID.FindStringSubmatch(youtubeURL)

	if len(matches) <= 1 {
		return "", fmt.Errorf("invalid youtube url: %s", youtubeURL)
	}

	return matches[1], nil
}
