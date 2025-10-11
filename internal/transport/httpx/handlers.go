package httpx

import (
	"errors"
	"go-quizz/m/internal/core/services/client"
	"go-quizz/m/internal/core/services/lobby"
	"go-quizz/m/internal/core/services/message"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Lobby   *lobby.Service
	Client  *client.Service
	Message *message.Service
}

func NewHandler(lobby *lobby.Service, client *client.Service, message *message.Service) *Handler {
	handler := &Handler{
		Router:  mux.NewRouter(),
		Lobby:   lobby,
		Client:  client,
		Message: message,
	}

	// routes
	handler.MapRoutes()

	return handler
}

func (h *Handler) MapRoutes() {
	// Lobby
	h.Router.HandleFunc("/api/lobbies", h.GetAllLobbies).Methods("GET")
	h.Router.HandleFunc("/api/lobby", h.PostLobby).Methods("POST")
	h.Router.HandleFunc("/api/lobby/{id}", h.GetLobby).Methods("GET")
	h.Router.HandleFunc("/api/lobby/{id}/clients", h.GetLobbyClients).Methods("GET")
	h.Router.HandleFunc("/api/lobby/{lobby_id}/connect/{client_id}", h.LobbyClientConnects).Methods("PATCH")
	h.Router.HandleFunc("/api/lobby/{lobby_id}/disconnect/{client_id}", h.LobbyClientDisconnects).Methods("PATCH")

	// Client
	h.Router.HandleFunc("/api/clients", h.GetAllClients).Methods("GET")
	h.Router.HandleFunc("/api/client", h.PostClient).Methods("POST")
	h.Router.HandleFunc("/api/client/{id}", h.GetClient).Methods("GET")

	// Message
	h.Router.HandleFunc("/api/messages", h.GetAllMessages).Methods("GET")
	h.Router.HandleFunc("/api/message", h.PostMessage).Methods("POST")
	h.Router.HandleFunc("/api/message/{id}", h.GetMessage).Methods("GET")
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
