package lobby

import (
	"errors"
	"fmt"
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

func (lobbySrvc *LobbyService) Generate() uuid.UUID {
	newLobbyID := uuid.New()
	if _, ok := lobbySrvc.lobbies[newLobbyID]; !ok {
		lobbySrvc.lobbies[newLobbyID] = domain.NewLobby(newLobbyID)
	} else {
		//TODO: handle error, lobby ID already exists
		fmt.Println("lobby already exists!")
	}

	return newLobbyID
}

func (lobbySrvc *LobbyService) GetAll() []domain.Lobby {
	var lobbies []domain.Lobby

	for _, lobby := range lobbySrvc.lobbies {
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

func (lobbySrvc *LobbyService) Get(lobbyID uuid.UUID) (domain.Lobby, error) {
	if _, ok := lobbySrvc.lobbies[lobbyID]; !ok {
		return domain.Lobby{}, errors.New("Lobby doesn't exists!")
	}

	return *lobbySrvc.lobbies[lobbyID], nil
}
