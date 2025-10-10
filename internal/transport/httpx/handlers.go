package httpx

import (
	"go-quizz/m/internal/core/services/client"
	"go-quizz/m/internal/core/services/lobby"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	Lobby  *lobby.Service
	Client *client.Service
}

func NewHandler(lobby *lobby.Service, client *client.Service) *Handler {
	handler := &Handler{
		Router: mux.NewRouter(),
		Lobby:  lobby,
		Client: client,
	}

	// routes
	handler.MapRoutes()

	return handler
}

func (h *Handler) MapRoutes() {
	// Lobby
	h.Router.HandleFunc("/api/lobbies", h.GetLobbies).Methods("GET")
	h.Router.HandleFunc("/api/lobby", h.PostLobby).Methods("POST")
	h.Router.HandleFunc("/api/lobby/{id}", h.GetLobby).Methods("GET")
	h.Router.HandleFunc("/api/lobby/{id}/clients", h.GetLobbyClients).Methods("GET")
	h.Router.HandleFunc("/api/lobby/{lobby_id}/connect/{client_id}", h.LobbyClientConnects).Methods("PATCH")
	h.Router.HandleFunc("/api/lobby/{lobby_id}/disconnect/{client_id}", h.LobbyClientDisconnects).Methods("PATCH")

	// Client
	h.Router.HandleFunc("/api/clients", h.GetClients).Methods("GET")
	h.Router.HandleFunc("/api/client", h.PostClient).Methods("POST")
	h.Router.HandleFunc("/api/client/{id}", h.GetClient).Methods("GET")
}
