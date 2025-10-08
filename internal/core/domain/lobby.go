package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Client struct {
	ID       uuid.UUID
	Username string
}

type Lobby struct {
	ID      uuid.UUID
	Clients map[uuid.UUID]*Client
	HostID  uuid.UUID
	Chat    *ChatRoom
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
		ID:      lobbyID,
		Clients: make(map[uuid.UUID]*Client, LobbyMaxClientCapacity),
		HostID:  hostID,
		Chat:    NewChatRoom(),
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

	//TODO: send message on the chat e.g. "User {username} connected."
}

func (lobby *Lobby) ClientDisconnect(clientID uuid.UUID) {
	if _, err := lobby.Clients[clientID]; err {
		//TODO: handle error client is not in lobby
		return
	}

	lobby.Clients[clientID] = nil
}
