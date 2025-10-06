package lobby

import (
	"fmt"
	"go-quizz/m/internal/core/domain"

	"github.com/google/uuid"
)

var Lobbies map[uuid.UUID]*domain.Lobby = make(map[uuid.UUID]*domain.Lobby)

func GenerateNewLobby(username string) *domain.Lobby {
	newLobbyID := uuid.New()
	if lobb, ok := Lobbies[newLobbyID]; ok {
		//TODO: handle error, lobby ID already exists
		fmt.Printf("lobby %v ID %v already exists!\n", lobb, newLobbyID)
		return nil
	}

	userID := uuid.New() // not important to check for duplicates when creating a lobby since its users array will be empty

	newLobby := domain.NewLobby(newLobbyID, userID, username)
	if newLobby != nil {
		Lobbies[newLobbyID] = newLobby
	}

	return newLobby
}
