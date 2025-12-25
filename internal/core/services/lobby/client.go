package lobby

import (
	"errors"
	"go-quizz/m/internal/core/domain"
	"strings"

	"github.com/google/uuid"
)

func (srvc *LobbyService) GetClientsInLobby(lobbyID uuid.UUID) ([]domain.Client, error) {
	if _, ok := srvc.lobbies[lobbyID]; !ok {
		return nil, errors.New("Lobby doesn't exist!")
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

func (srvc *LobbyService) GetClientInLobby(lobbyID uuid.UUID, clientID uuid.UUID) (domain.Client, error) {
	if _, ok := srvc.lobbies[lobbyID]; !ok {
		return domain.Client{}, errors.New("Lobby doesn't exist!")
	}

	lobby := srvc.lobbies[lobbyID]

	client, ok := lobby.Clients[clientID]
	if !ok {
		return domain.Client{}, errors.New("client is not in lobby")
	}

	clientCopy := domain.Client{
		ID:       client.ID,
		Username: client.Username,
	}

	return clientCopy, nil
}

func (srvc *LobbyService) ConnectsClient(lobbyID uuid.UUID, username string) (domain.Client, error) {
	lobby, ok := srvc.lobbies[lobbyID]
	if !ok {
		return domain.Client{}, errors.New("lobby doesn't exist")
	}
	if len(lobby.Clients) == domain.LobbyMaxClientCapacity {
		return domain.Client{}, errors.New("lobby is already full")
	}
	if strings.TrimSpace(username) == "" {
		return domain.Client{}, errors.New("username is not valid")
	}

	clientID := uuid.New()

	client := domain.NewClient(clientID, username)
	lobby.Clients[clientID] = client

	if len(lobby.Clients) == 1 {
		lobby.HostID = client.ID
	}

	return *client, nil
}

func (srvc *LobbyService) DisconnectsClient(lobbyID uuid.UUID, client domain.Client) (domain.Lobby, error) {
	lobby, ok := srvc.lobbies[lobbyID]
	if !ok {
		return domain.Lobby{}, errors.New("lobby doesn't exist")
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
