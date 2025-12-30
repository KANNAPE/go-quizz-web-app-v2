package handlers

import (
	"encoding/json"
	"fmt"
	"go-quizz/m/internal/transport/api/dto"
	"net/http"
)

func (h *Handler) GetLobbyClients(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.GetLobbyClientsResponse]()

	// invalid lobby ID
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("get lobby clients[%v]: %w", lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// lobby not found
	lobby, err := h.Lobby.GetLobby(lobbyID)
	if err != nil {
		apiResponse.Code = http.StatusNotFound
		apiResponse.Message = fmt.Errorf("get lobby clients[%v]: %w", lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	getLobbyClientsResponse := dto.GetLobbyClientsResponse{
		LobbyID: lobbyID,
	}

	// for every client in the lobby, we append a new value in our response
	for _, client := range lobby.Clients {
		getLobbyClientsResponse.Clients = append(getLobbyClientsResponse.Clients, dto.GetClientResponse{
			ID:       client.ID,
			Username: client.Username,
		})
	}

	apiResponse.Data = getLobbyClientsResponse

	// sending response
	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}

func (h *Handler) LobbyClientConnects(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.LobbyClientConnectsResponse]()

	// retreiving the body of the request
	var clientConnectionReq dto.ClientConnectionRequest
	if err := json.NewDecoder(req.Body).Decode(&clientConnectionReq); err != nil {
		apiResponse.Code = http.StatusUnprocessableEntity
		apiResponse.Message = fmt.Errorf("client lobby connection: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}
		return
	}

	// lobby ID is invalid
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("client lobby connection: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// lobby doesn't exist
	client, err := h.Lobby.ConnectsClient(lobbyID, clientConnectionReq.Username)
	if err != nil {
		apiResponse.Code = http.StatusBadRequest // todo: handle different error codes
		apiResponse.Message = fmt.Errorf("client lobby connection[%v]: %w", lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// sending response
	clientConnectionResponse := dto.LobbyClientConnectsResponse{
		LobbyID:  lobbyID,
		ClientID: client.ID,
		Username: client.Username,
	}

	apiResponse.Data = clientConnectionResponse

	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}

func (h *Handler) LobbyClientDisconnects(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.LobbyClientDisconnectsResponse]()

	// invalid lobby ID
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("client lobby disconnection: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// invalid client ID
	clientID, err := getUUIDFromUri(req, "client_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("client lobby disconnection: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// lobby or client doesn't exist
	client, err := h.Lobby.GetClientInLobby(lobbyID, clientID)
	if err != nil {
		apiResponse.Code = http.StatusNotFound
		apiResponse.Message = fmt.Errorf("client lobby disconnection: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	h.Lobby.DisconnectsClient(lobbyID, client)

	// no data

	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}
