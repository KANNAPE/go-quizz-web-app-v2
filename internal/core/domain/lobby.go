package domain

import (
	"github.com/google/uuid"
)

type Lobby struct {
	ID       uuid.UUID
	HostID   uuid.UUID
	Clients  map[uuid.UUID]*Client
	Messages map[uuid.UUID]*Message
}

var LobbyMaxClientCapacity int = 8

// functions
func NewLobby(lobbyID uuid.UUID) *Lobby {
	return &Lobby{
		ID:       lobbyID,
		Clients:  make(map[uuid.UUID]*Client),
		Messages: make(map[uuid.UUID]*Message),
	}
}
