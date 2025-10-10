package httpx

import (
	"go-quizz/m/internal/core/services/lobby"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	Lobby  *lobby.LobbyService
}

func NewHandler(lobby *lobby.LobbyService) *Handler {
	handler := &Handler{
		Lobby:  lobby,
		Router: mux.NewRouter(),
	}

	// routes
	handler.MapRoutes()

	return handler
}

func (h *Handler) MapRoutes() {
	h.Router.HandleFunc("/api/lobbies", h.GetLobbies).Methods("GET")
	h.Router.HandleFunc("/api/lobby", h.PostLobby).Methods("POST")
	h.Router.HandleFunc("/api/lobby/{id}", h.GetLobby).Methods("GET")
}
