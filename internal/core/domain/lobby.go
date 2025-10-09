package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Lobby struct {
	ID       uuid.UUID              `json:"id"`
	Clients  map[uuid.UUID]*Client  `json:"clients"`
	HostID   uuid.UUID              `json:"host_id"`
	Messages map[uuid.UUID]*Message `json:"chat"`
}

var LobbyMaxClientCapacity int = 8

// functions
func NewLobby(lobbyID uuid.UUID, hostID uuid.UUID, hostUsername string) *Lobby {
	if hostUsername == "" {
		// TODO: handle error
		fmt.Println("host username is empty")
		return nil
	}

	newLobby := &Lobby{
		ID:       lobbyID,
		Clients:  make(map[uuid.UUID]*Client, LobbyMaxClientCapacity),
		HostID:   hostID,
		Messages: make(map[uuid.UUID]*Message, 0),
	}

	newLobby.ClientConnect(hostID, hostUsername)

	return newLobby
}

func (lobby *Lobby) ClientConnect(clientID uuid.UUID, clientUsername string) {
	if clientUsername == "" {
		//TODO: handle error invalid username
		return
	}
	if len(lobby.Clients) == LobbyMaxClientCapacity {
		//TODO: handle error lobby is full
		return
	}
	if _, ok := lobby.Clients[clientID]; ok {
		//TODO: handle error client has already joined
		return
	}

	newClient := &Client{ID: clientID, Username: clientUsername}

	lobby.Clients[clientID] = newClient
}

func (lobby *Lobby) ClientDisconnect(clientID uuid.UUID) {
	if _, ok := lobby.Clients[clientID]; !ok {
		//TODO: handle error client is not in lobby
		return
	}

	delete(lobby.Clients, clientID)
}
