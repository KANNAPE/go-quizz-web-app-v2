package transport

import "go-quizz/m/internal/core/domain"

type LobbyService interface {
	GetLobby() (domain.Lobby, error)
	PostLobby() (domain.Lobby, error)
	UpdateLobby() (domain.Lobby, error)
	DeleteLobby() error
}
