package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID `json:"id"`
	SenderID uuid.UUID `json:"sender_id"`
	TimeSent time.Time `json:"time_sent"`
	Body     string    `json:"content"`
}

func NewMessage(id uuid.UUID, timeSent time.Time, body string) *Message {
	return &Message{
		ID:       id,
		TimeSent: timeSent,
		Body:     body,
	}
}
