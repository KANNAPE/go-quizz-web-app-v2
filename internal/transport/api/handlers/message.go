package handlers

import (
	"encoding/json"
	"fmt"
	"go-quizz/m/internal/transport/api/dto"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) GetLobbyMessages(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.GetLobbyMessagesResponse]()

	// invalid lobby ID
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("get lobby messages: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// lobby doesn't exist
	messages, err := h.Lobby.GetAllMessagesInLobby(lobbyID)
	if err != nil {
		apiResponse.Code = http.StatusNotFound
		apiResponse.Message = fmt.Errorf("get lobby messages[%v]: %w", lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// creating the response for every message in the lobby
	getLobbyMessagesResponse := dto.GetLobbyMessagesResponse{
		LobbyID: lobbyID,
	}

	for _, message := range messages {
		getLobbyMessagesResponse.Messages = append(getLobbyMessagesResponse.Messages, dto.GetMessageResponse{
			ID:       message.ID,
			SenderID: message.SenderID,
			TimeSent: message.TimeSent,
			Body:     message.Body,
		})
	}

	apiResponse.Data = getLobbyMessagesResponse

	// sending the response
	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}

func (h *Handler) GetMessage(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.GetLobbyMessageResponse]()

	// invalid lobby ID
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("get message in lobby: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// invalid message ID
	messageID, err := getUUIDFromUri(req, "message_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("get message in lobby: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// message or lobby doesn't exist
	message, err := h.Lobby.GetLobbyMessage(lobbyID, messageID)
	if err != nil {
		apiResponse.Code = http.StatusNotFound
		apiResponse.Message = fmt.Errorf("get message[%v] in lobby[%v]: %w", messageID, lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// message found
	apiResponse.Data = dto.GetLobbyMessageResponse{
		LobbyID: lobbyID,
		Message: dto.GetMessageResponse{
			ID:       message.ID,
			SenderID: message.SenderID,
			TimeSent: message.TimeSent,
			Body:     message.Body,
		},
	}

	// sending the response
	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}

func (h *Handler) PostMessage(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.PostMessageResponse]()

	// error in the request
	var messageReq dto.PostMessageRequest
	if err := json.NewDecoder(req.Body).Decode(&messageReq); err != nil {
		apiResponse.Code = http.StatusUnprocessableEntity
		apiResponse.Message = fmt.Errorf("post lobby message: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// invalid lobby ID
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("post lobby message: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// invalid ID in request body
	senderID, err := uuid.Parse(messageReq.SenderID)
	if err != nil {
		apiResponse.Code = http.StatusUnprocessableEntity
		apiResponse.Message = fmt.Errorf("post lobby message: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// error in message creation
	messageID, err := h.Lobby.CreateMessage(lobbyID, senderID, messageReq.Body)
	if err != nil {
		apiResponse.Code = http.StatusInternalServerError
		apiResponse.Message = fmt.Errorf("post lobby message: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// message created
	apiResponse.Data = dto.PostMessageResponse{
		ID:       messageID,
		SenderID: senderID,
	}

	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}
