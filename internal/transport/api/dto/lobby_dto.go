package dto

import (
	"go-quizz/m/internal/core/domain"

	"github.com/google/uuid"
)

type GetLobbyResponse struct {
	ID       uuid.UUID                     `json:"id"`
	HostID   uuid.UUID                     `json:"host_id"`
	Clients  map[uuid.UUID]*domain.Client  `json:"clients"`
	Messages map[uuid.UUID]*domain.Message `json:"messages"`
}

type GetAllLobbiesResponse struct {
	Lobbies []GetLobbyResponse `json:"lobbies"`
}

type CreateLobbyResponse struct {
	ID uuid.UUID `json:"created_lobby_id"`
}

type DeleteLobbyResponse struct {
}
