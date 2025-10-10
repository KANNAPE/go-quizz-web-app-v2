package domain

import (
	"errors"

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
func NewLobby(lobbyID uuid.UUID, hostID uuid.UUID, hostUsername string) (*Lobby, error) {
	if hostUsername == "" {
		// TODO: handle error
		return nil, errors.New("host username is empty")
	}

	newLobby := &Lobby{
		ID:       lobbyID,
		Clients:  make(map[uuid.UUID]*Client, LobbyMaxClientCapacity),
		HostID:   hostID,
		Messages: make(map[uuid.UUID]*Message, 0),
	}

	if err := newLobby.ClientConnects(hostID, hostUsername); err != nil {
		return nil, errors.New("could not connect client")
	}
	return newLobby, nil
}

func (lobby *Lobby) ClientConnects(clientID uuid.UUID, clientUsername string) error {
	if clientUsername == "" {
		//TODO: handle error
		return errors.New("invalid username")
	}
	if len(lobby.Clients) == LobbyMaxClientCapacity {
		//TODO: handle error
		return errors.New("lobby is full")
	}
	if _, ok := lobby.Clients[clientID]; ok {
		//TODO: handle error
		return errors.New("client has already joined")
	}

	newClient := NewClient(clientID, clientUsername)

	lobby.Clients[clientID] = newClient

	return nil
}

func (lobby *Lobby) ClientDisconnects(clientID uuid.UUID) error {
	if _, ok := lobby.Clients[clientID]; !ok {
		//TODO: handle error
		return errors.New("client is not in lobby")
	}

	delete(lobby.Clients, clientID)

	return nil
}

func (lobby *Lobby) ClientSendsMessage(clientID uuid.UUID, messageBody string) error {
	return nil
}
