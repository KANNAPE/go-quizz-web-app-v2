package handlers

import (
	"encoding/json"
	"errors"
	"go-quizz/m/internal/core/services/lobby"
	"go-quizz/m/internal/transport/api/dto"
	"go-quizz/m/internal/transport/api/middlewares"
	"go-quizz/m/internal/transport/websocket"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	Router *http.ServeMux
	Lobby  *lobby.LobbyService
	HubMgr *websocket.HubManager
}

func NewHandler(lobby *lobby.LobbyService, hubMgr *websocket.HubManager) *Handler {
	handler := &Handler{
		Router: http.NewServeMux(),
		Lobby:  lobby,
		HubMgr: hubMgr,
	}

	// routes
	handler.MapRoutes()

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Apply middleware to the router
	middlewares.CORSMiddleware(h.Router).ServeHTTP(w, r)
}

func (h *Handler) MapRoutes() {
	// Lobby
	h.Router.HandleFunc("GET /api/lobbies", h.GetAllLobbies)
	h.Router.HandleFunc("POST /api/lobby", h.PostLobby)
	h.Router.HandleFunc("GET /api/lobby/{lobby_id}", h.GetLobby)
	h.Router.HandleFunc("DELETE /api/lobby/{lobby_id}", h.DeleteLobby)

	// Client
	h.Router.HandleFunc("GET /api/lobby/{lobby_id}/clients", h.GetLobbyClients)
	h.Router.HandleFunc("POST /api/lobby/{lobby_id}/connect", h.LobbyClientConnects)
	h.Router.HandleFunc("DELETE /api/lobby/{lobby_id}/disconnect/{client_id}", h.LobbyClientDisconnects)

	// Message
	h.Router.HandleFunc("GET /api/lobby/{lobby_id}/messages", h.GetLobbyMessages)
	h.Router.HandleFunc("POST /api/lobby/{lobby_id}/message", h.PostMessage)
	h.Router.HandleFunc("GET /api/lobby/{lobby_id}/message/{message_id}", h.GetMessage)

	// Websocket
	h.Router.HandleFunc("GET /api/lobby/{lobby_id}/ws", h.ServeWebSocket)
}

func getUUIDFromUri(req *http.Request, uriID string) (uuid.UUID, error) {
	stringID := req.PathValue(uriID)
	if stringID == "" {
		return uuid.Nil, errors.New("invalid uri ID")
	}

	ID, err := uuid.Parse(stringID)
	if err != nil {
		return uuid.Nil, errors.New("ID is not a valid UUID")
	}

	return ID, nil
}

func encodeResponse[T any](writer http.ResponseWriter, response dto.APIResponse[T]) error {
	writer.WriteHeader(response.Code)
	return json.NewEncoder(writer).Encode(response)
}
