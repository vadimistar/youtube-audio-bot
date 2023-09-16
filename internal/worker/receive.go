package worker

import (
	"encoding/json"
	"io"
)

func receiveMessages(body io.Reader) (msgs messages, err error) {
	err = json.NewDecoder(body).Decode(&msgs)
	if err != nil {
		return messages{}, err
	}
	return msgs, nil
}

type messages struct {
	Messages []struct {
		Details struct {
			Message struct {
				Body string `json:"body"`
			} `json:"message"`
		} `json:"details"`
	} `json:"messages"`
}
