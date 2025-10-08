package handlers

import "go-quizz/m/internal/httpx/dto"

type Handler struct {
	JoinURL string
	Lobby   dto.Lobby
	Host    dto.Client
}

func NewHandler() *Handler {
	handler := &Handler{
		JoinURL: "",
		Lobby:   dto.Lobby{},
		Host:    dto.Client{},
	}

	return handler
}
