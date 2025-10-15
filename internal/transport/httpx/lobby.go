package httpx

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllLobbies(writer http.ResponseWriter, req *http.Request) {
	lobbies := h.Lobby.GetAllLobbies()

	if len(lobbies) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(writer).Encode(lobbies); err != nil {
		panic(err)
	}
}

func (h *Handler) GetLobby(writer http.ResponseWriter, req *http.Request) {
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, err := h.Lobby.GetLobby(lobbyID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(writer).Encode(lobby); err != nil {
		panic(err)
	}
}

func (h *Handler) PostLobby(writer http.ResponseWriter, req *http.Request) {
	lobby_id := h.Lobby.OpenLobby()

	if err := json.NewEncoder(writer).Encode(lobby_id); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteLobby(writer http.ResponseWriter, req *http.Request) {
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Lobby.CloseLobby(lobbyID); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
