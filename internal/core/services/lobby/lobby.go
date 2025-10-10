package lobby

import (
	"fmt"
	"go-quizz/m/internal/core/domain"

	"github.com/google/uuid"
)

type LobbyService struct {
	lobbies map[uuid.UUID]*domain.Lobby
}

func NewService() *LobbyService {
	return &LobbyService{}
}

func (lobbySrvc *LobbyService) Generate(username string) string {
	newLobbyID := uuid.New()
	if _, ok := lobbySrvc.lobbies[newLobbyID]; ok {
		//TODO: handle error, lobby ID already exists
		fmt.Println("lobby already exists!")

		return ""
	}

	userID := uuid.New() // not important to check for duplicates when creating a lobby since its users array will be empty

	if newLobby, error := domain.NewLobby(newLobbyID, userID, username); error != nil {
		lobbySrvc.lobbies[newLobbyID] = newLobby
	}

	return "/api/lobby/" + newLobbyID.String()
}
