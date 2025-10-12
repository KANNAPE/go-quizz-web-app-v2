package httpx

import (
	"errors"
	service "go-quizz/m/internal/core/service/lobby"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	Lobby  *service.Lobby
}

func NewHandler(lobby *service.Lobby) *Handler {
	handler := &Handler{
		Router: mux.NewRouter(),
		Lobby:  lobby,
	}

	// routes
	handler.MapRoutes()

	return handler
}

func (h *Handler) MapRoutes() {
	// Lobby
	h.Router.HandleFunc("/api/lobbies", h.GetAllLobbies).Methods("GET")
	h.Router.HandleFunc("/api/lobby", h.PostLobby).Methods("POST")
	h.Router.HandleFunc("/api/lobby/{lobby_id}", h.GetLobby).Methods("GET")
	h.Router.HandleFunc("/api/lobby/{lobby_id}", h.DeleteLobby).Methods("DELETE")

	// Client
	h.Router.HandleFunc("/api/lobby/{lobby_id}/clients", h.GetLobbyClients).Methods("GET")
	h.Router.HandleFunc("/api/lobby/{lobby_id}/connect", h.LobbyClientConnects).Methods("POST")
	h.Router.HandleFunc("/api/lobby/{lobby_id}/disconnect/{client_id}", h.LobbyClientDisconnects).Methods("DELETE")

	// Message
	h.Router.HandleFunc("/api/lobby/{lobby_id}/messages", h.GetLobbyMessages).Methods("GET")
	h.Router.HandleFunc("/api/lobby/{lobby_id}/message", h.PostMessage).Methods("POST")
	h.Router.HandleFunc("/api/lobby/{lobby_id}/message/{message_id}", h.GetMessage).Methods("GET")
}

func getUUIDFromUri(req *http.Request, uriID string) (uuid.UUID, error) {
	vars := mux.Vars(req)

	stringID := vars[uriID]
	if stringID == "" {
		return uuid.Nil, errors.New("invalid uri ID")
	}

	ID, err := uuid.Parse(stringID)
	if err != nil {
		return uuid.Nil, errors.New("ID is not a valid UUID")
	}

	return ID, nil
}
