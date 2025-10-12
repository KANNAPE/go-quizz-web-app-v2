package lobby

import (
	"errors"
	"go-quizz/m/internal/core/domain"
	"strings"

	"github.com/google/uuid"
)

func (srvc *Service) GetClientsInLobby(lobbyID uuid.UUID) ([]domain.Client, error) {
	if _, ok := srvc.lobbies[lobbyID]; !ok {
		return nil, errors.New("Lobby doesn't exists!")
	}

	lobby := srvc.lobbies[lobbyID]

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

func (srvc *Service) ConnectsClient(lobbyID uuid.UUID, username string) (domain.Lobby, error) {
	lobby, ok := srvc.lobbies[lobbyID]
	if !ok {
		return domain.Lobby{}, errors.New("lobby doesn't exists")
	}
	if len(lobby.Clients) == domain.LobbyMaxClientCapacity {
		return domain.Lobby{}, errors.New("lobby is already full")
	}
	if strings.Trim(username, "") == "" {
		return domain.Lobby{}, errors.New("username is not valid")
	}

	clientID := uuid.New()

	client := domain.NewClient(clientID, username)
	lobby.Clients[clientID] = client

	if len(lobby.Clients) == 1 {
		lobby.HostID = client.ID
	}

	return *lobby, nil
}

func (srvc *Service) DisconnectsClient(lobbyID uuid.UUID, client domain.Client) (domain.Lobby, error) {
	lobby, ok := srvc.lobbies[lobbyID]
	if !ok {
		return domain.Lobby{}, errors.New("lobby doesn't exists")
	}
	if len(lobby.Clients) == 0 {
		return domain.Lobby{}, errors.New("lobby is already empty")
	}
	if _, ok := lobby.Clients[client.ID]; !ok {
		return domain.Lobby{}, errors.New("client is not in lobby")
	}

	delete(lobby.Clients, client.ID)

	if lobby.HostID == client.ID {
		if err := srvc.CloseLobby(lobbyID); err != nil {
			return domain.Lobby{}, errors.New("could not close lobby")
		}

		return domain.Lobby{}, nil
	}

	return *lobby, nil
}
