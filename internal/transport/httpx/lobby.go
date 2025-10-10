package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) GetLobbies(writer http.ResponseWriter, req *http.Request) {
	lobbies := h.Lobby.GetAll()

	if len(lobbies) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(writer).Encode(lobbies); err != nil {
		panic(err)
	}
}

func (h *Handler) GetLobby(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	lobby_id := vars["id"]
	if lobby_id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	lobbyID, err := uuid.Parse(lobby_id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, err := h.Lobby.Get(lobbyID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(writer).Encode(lobby); err != nil {
		panic(err)
	}
}

func (h *Handler) PostLobby(writer http.ResponseWriter, req *http.Request) {
	lobby_id := h.Lobby.Generate()

	if err := json.NewEncoder(writer).Encode(lobby_id); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
