package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID
	SenderID uuid.UUID
	TimeSent time.Time
	Body     string
}

func NewMessage(messageID uuid.UUID, senderID uuid.UUID, timeSent time.Time, body string) *Message {
	return &Message{
		ID:       messageID,
		SenderID: senderID,
		TimeSent: timeSent,
		Body:     body,
	}
}
