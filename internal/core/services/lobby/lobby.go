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

func (lobbySrvc *Service) Generate() uuid.UUID {
	newLobbyID := uuid.New()
	if _, ok := lobbySrvc.lobbies[newLobbyID]; !ok {
		lobbySrvc.lobbies[newLobbyID] = domain.NewLobby(newLobbyID)
	} else {
		//TODO: handle error, lobby ID already exists
		fmt.Println("lobby already exists!")
	}

	return newLobbyID
}

func (lobbySrvc *Service) GetAll() []domain.Lobby {
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

func (lobbySrvc *Service) Get(lobbyID uuid.UUID) (domain.Lobby, error) {
	if _, ok := lobbySrvc.lobbies[lobbyID]; !ok {
		return domain.Lobby{}, errors.New("Lobby doesn't exists!")
	}

	return *lobbySrvc.lobbies[lobbyID], nil
}

func (lobbySrvc *Service) GetClients(lobbyID uuid.UUID) ([]domain.Client, error) {
	if _, ok := lobbySrvc.lobbies[lobbyID]; !ok {
		return nil, errors.New("Lobby doesn't exists!")
	}

	lobby := lobbySrvc.lobbies[lobbyID]

	var clients []domain.Client
	for _, client := range lobby.Clients {
		clientCopy := domain.Client{
			ID:       client.ID,
			Username: client.Username,
		}

		clients = append(clients, clientCopy)
	}

	return clients, nil
}

func (lobbySrvc *Service) ConnectsClient(lobbyID uuid.UUID, client domain.Client) error {
	if lobby, ok := lobbySrvc.lobbies[lobbyID]; !ok {
		return errors.New("lobby doesn't exists")
	} else if len(lobby.Clients) == domain.LobbyMaxClientCapacity {
		return errors.New("lobby is already full")
	} else if _, ok := lobby.Clients[client.ID]; ok {
		return errors.New("client already in lobby")
	} else {
		lobby.Clients[client.ID] = &client

		if len(lobby.Clients) == 1 {
			lobby.HostID = client.ID
		}
		return nil
	}
}

func (lobbySrvc *Service) DisconnectsClient(lobbyID uuid.UUID, client domain.Client) error {
	if lobby, ok := lobbySrvc.lobbies[lobbyID]; !ok {
		return errors.New("lobby doesn't exists")
	} else if len(lobby.Clients) == 0 {
		return errors.New("lobby is already empty")
	} else if _, ok := lobby.Clients[client.ID]; !ok {
		return errors.New("client is not in lobby")
	} else {
		delete(lobby.Clients, client.ID)

		if lobby.HostID == client.ID {
			// TODO
		}
		return nil
	}
}
