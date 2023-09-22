package entity

type TaskRequest struct {
	ChatID  int64  `json:"chat_id"`
	VideoID string `json:"video_id"`
}

type TaskResponse struct {
	ChatID int64  `json:"chat_id"`
	Key    string `json:"key,omitempty"`
	Error  string `json:"error,omitempty"`
}
