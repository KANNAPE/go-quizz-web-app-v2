package domain

import (
	"github.com/google/uuid"
)

type Lobby struct {
	ID       uuid.UUID              `json:"id"`
	HostID   uuid.UUID              `json:"host_id"`
	Clients  map[uuid.UUID]*Client  `json:"clients"`
	Messages map[uuid.UUID]*Message `json:"messages"`
}

var LobbyMaxClientCapacity int = 8

// functions
func NewLobby(lobbyID uuid.UUID) *Lobby {
	return &Lobby{
		ID:      lobbyID,
		Clients: make(map[uuid.UUID]*Client, LobbyMaxClientCapacity),
	}
}
