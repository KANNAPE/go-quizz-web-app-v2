package transport

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Server  *http.Server
	Lobby   LobbyService
	Message MessageService
}

func NewHandler() *Handler {
	handler := &Handler{}

	return handler
}
