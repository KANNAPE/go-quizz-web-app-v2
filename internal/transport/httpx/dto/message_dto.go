package dto

import (
	"time"

	"github.com/google/uuid"
)

// Requests
type PostMessageRequest struct {
	SenderID string `json:"sender_id" validate:"required"`
	Body     string `json:"body" validate:"required"`
}

// Responses
type GetMessageResponse struct {
	ID       uuid.UUID `json:"id"`
	SenderID uuid.UUID `json:"sender_id"`
	TimeSent time.Time `json:"time_sent"`
	Body     string    `json:"body"`
}

type GetLobbyMessagesResponse struct {
	LobbyID  uuid.UUID            `json:"lobby_id"`
	Messages []GetMessageResponse `json:"messages"`
}

type GetLobbyMessageResponse struct {
	LobbyID uuid.UUID          `json:"lobby_id"`
	Message GetMessageResponse `json:"message"`
}

type PostMessageResponse struct {
	ID       uuid.UUID `json:"id"`
	SenderID uuid.UUID `json:"sender_id"`
}
