package lobby

import (
	"errors"
	"fmt"
	"go-quizz/m/internal/core/domain"

	"github.com/google/uuid"
)

type Service struct {
	lobbies map[uuid.UUID]*domain.Lobby
}

func NewService() *Service {
	return &Service{
		lobbies: make(map[uuid.UUID]*domain.Lobby),
	}
}

func (srvc *Service) OpenLobby() uuid.UUID {
	newLobbyID := uuid.New()
	if _, ok := srvc.lobbies[newLobbyID]; !ok {
		srvc.lobbies[newLobbyID] = domain.NewLobby(newLobbyID)
	} else {
		//TODO: handle error, lobby ID already exists
		fmt.Println("lobby already exists!")
	}

	return newLobbyID
}

func (srvc *Service) GetAllLobbies() []domain.Lobby {
	var lobbies []domain.Lobby

	for _, lobby := range srvc.lobbies {
		lobbyCopy := domain.Lobby{
			ID:       lobby.ID,
			HostID:   lobby.HostID,
			Clients:  lobby.Clients,
			Messages: lobby.Messages,
		}

		lobbies = append(lobbies, lobbyCopy)
	}

	return lobbies
}

func (srvc *Service) GetLobby(lobbyID uuid.UUID) (domain.Lobby, error) {
	if _, ok := srvc.lobbies[lobbyID]; !ok {
		return domain.Lobby{}, errors.New("Lobby doesn't exist!")
	}

	return *srvc.lobbies[lobbyID], nil
}

func (srvc *Service) CloseLobby(lobbyID uuid.UUID) error {
	lobby, ok := srvc.lobbies[lobbyID]
	if !ok {
		return errors.New("lobby doesn't exist!")
	}

	for clientID := range lobby.Clients {
		delete(lobby.Clients, clientID)
	}
	for messageID := range lobby.Messages {
		delete(lobby.Messages, messageID)
	}

	delete(srvc.lobbies, lobbyID)

	return nil
}
