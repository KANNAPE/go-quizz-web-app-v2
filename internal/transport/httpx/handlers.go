package httpx

import (
	"fmt"
	"go-quizz/m/internal/core/services/lobby"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	Lobby  *lobby.LobbyService
}

func NewHandler(lobby *lobby.LobbyService) *Handler {
	handler := &Handler{
		Lobby: lobby,
	}

	handler.Router = mux.NewRouter()

	// routes
	handler.Router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, "API Endpoint Hit")
	})

	return handler
}
