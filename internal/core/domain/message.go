package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID       uuid.UUID `json:"id"`
	Sender   *Client   `json:"sender"`
	TimeSent time.Time `json:"time_sent"`
	Body     string    `json:"content"`
}
