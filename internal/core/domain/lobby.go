package domain

import (
	"github.com/google/uuid"
)

type Client struct {
	Username string
}

type Lobby struct {
	ID      uuid.UUID
	Clients []*Client
	Chat    *ChatRoom
}

var LobbyMaxClientCapacity int = 8

// functions
func CreateLobby(ID uuid.UUID, hostUsername string) *Lobby {
	if hostUsername == "" {
		// TODO: handle error
		return nil
	}

	newLobby := &Lobby{
		ID:      ID,
		Clients: make([]*Client, LobbyMaxClientCapacity),
	}

	host := &Client{Username: hostUsername}

	newLobby.Clients[0] = host

	return newLobby
}

func (lobby *Lobby) AddClient(username string) {
	if username == "" {
		//TODO: handle error
		return
	}
	if len(lobby.Clients) == LobbyMaxClientCapacity {
		//TODO: handle error
		return
	}

	newClient := &Client{Username: username}

	for slot, client := range lobby.Clients {
		if client == nil {
			lobby.Clients[slot] = newClient
			break
		}
	}
}
