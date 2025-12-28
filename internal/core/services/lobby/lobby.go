package lobby

import (
	"errors"
	"go-quizz/m/internal/core/domain"
	"sync"

	"github.com/google/uuid"
)

type LobbyService struct {
	mu      sync.RWMutex
	lobbies map[uuid.UUID]*domain.Lobby
}

func NewService() *LobbyService {
	return &LobbyService{
		lobbies: make(map[uuid.UUID]*domain.Lobby),
	}
}

func (srvc *LobbyService) OpenLobby() (uuid.UUID, error) {
	srvc.mu.Lock()
	defer srvc.mu.Unlock()

	newLobbyID := uuid.New()
	if _, ok := srvc.lobbies[newLobbyID]; !ok {
		srvc.lobbies[newLobbyID] = domain.NewLobby(newLobbyID)
	} else {
		return uuid.UUID{}, errors.New("lobby already exists!")
	}

	return newLobbyID, nil
}

func (srvc *LobbyService) GetAllLobbies() []domain.Lobby {
	srvc.mu.RLock()
	defer srvc.mu.RUnlock()

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
	srvc.mu.RLock()
	defer srvc.mu.RUnlock()

	if _, ok := srvc.lobbies[lobbyID]; !ok {
		return domain.Lobby{}, errors.New("Lobby doesn't exist!")
	}

	return *srvc.lobbies[lobbyID], nil
}

func (srvc *LobbyService) CloseLobby(lobbyID uuid.UUID) error {
	srvc.mu.Lock()
	defer srvc.mu.Unlock()

	return srvc.closeLobby(lobbyID)
}

func (srvc *LobbyService) closeLobby(lobbyID uuid.UUID) error {
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
