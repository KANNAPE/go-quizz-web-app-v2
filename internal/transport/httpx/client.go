package httpx

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetLobbyClients(writer http.ResponseWriter, req *http.Request) {
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

	if err := json.NewEncoder(writer).Encode(lobby.Clients); err != nil {
		panic(err)
	}
}

type clientConnectionRequest struct {
	username string `validate:"required"`
}

func (h *Handler) LobbyClientConnects(writer http.ResponseWriter, req *http.Request) {
	var clientConnectionReq clientConnectionRequest
	if err := json.NewDecoder(req.Body).Decode(&clientConnectionReq); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, err := h.Lobby.ConnectsClient(lobbyID, clientConnectionReq.username)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(writer).Encode(lobby); err != nil {
		panic(err)
	}
}

func (h *Handler) LobbyClientDisconnects(writer http.ResponseWriter, req *http.Request) {
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	clientID, err := getUUIDFromUri(req, "client_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	client, err := h.Lobby.GetClientInLobby(lobbyID, clientID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	lobby, _ := h.Lobby.DisconnectsClient(lobbyID, client)

	if err := json.NewEncoder(writer).Encode(lobby); err != nil {
		panic(err)
	}
}
