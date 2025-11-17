package dto

import "time"

type Question struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type QuestionRequest struct {
	Text string `json:"text"`
}
