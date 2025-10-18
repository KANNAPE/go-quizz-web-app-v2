package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) GetLobbyMessages(writer http.ResponseWriter, req *http.Request) {
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	messages, err := h.Lobby.GetAllMessagesInLobby(lobbyID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(writer).Encode(messages); err != nil {
		panic(err)
	}
}

func (h *Handler) GetMessage(writer http.ResponseWriter, req *http.Request) {
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	messageID, err := getUUIDFromUri(req, "message_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	message, err := h.Lobby.GetLobbyMessage(lobbyID, messageID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(writer).Encode(message); err != nil {
		panic(err)
	}
}

type PostMessageRequest struct {
	Body     string `json:"content" validate:"required"`
	SenderID string `json:"sender_id" validate:"required"`
}

func (h *Handler) PostMessage(writer http.ResponseWriter, req *http.Request) {
	var messageReq PostMessageRequest
	if err := json.NewDecoder(req.Body).Decode(&messageReq); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	senderID, err := uuid.Parse(messageReq.SenderID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	messageID, err := h.Lobby.CreateMessage(lobbyID, senderID, messageReq.Body)
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := json.NewEncoder(writer).Encode(messageID); err != nil {
		panic(err)
	}
}
