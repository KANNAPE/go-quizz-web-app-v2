package lobby

import (
	"errors"
	"go-quizz/m/internal/core/domain"

	"github.com/google/uuid"
)

type LobbyService struct {
	lobbies map[uuid.UUID]*domain.Lobby
}

func NewService() *LobbyService {
	return &LobbyService{
		lobbies: make(map[uuid.UUID]*domain.Lobby),
	}
}

func (srvc *LobbyService) OpenLobby() (uuid.UUID, error) {
	newLobbyID := uuid.New()
	if _, ok := srvc.lobbies[newLobbyID]; !ok {
		srvc.lobbies[newLobbyID] = domain.NewLobby(newLobbyID)
	} else {
		return uuid.UUID{}, errors.New("lobby already exists!")
	}

	return newLobbyID, nil
}

func (srvc *LobbyService) GetAllLobbies() []domain.Lobby {
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

func (srvc *LobbyService) GetLobby(lobbyID uuid.UUID) (domain.Lobby, error) {
	if _, ok := srvc.lobbies[lobbyID]; !ok {
		return domain.Lobby{}, errors.New("Lobby doesn't exist!")
	}

	return *srvc.lobbies[lobbyID], nil
}

func (srvc *LobbyService) CloseLobby(lobbyID uuid.UUID) error {
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
