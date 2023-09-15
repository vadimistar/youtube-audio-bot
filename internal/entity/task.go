package entity

type TaskRequest struct {
	ChatID   int64  `json:"chat_id"`
	VideoURL string `json:"video_url"`
}

type TaskResponse struct {
	ChatID       int64  `json:"chat_id"`
	FileLocation string `json:"file_location,omitempty"`
	Error        string `json:"error,omitempty"`
}
