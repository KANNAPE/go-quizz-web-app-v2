package handlers

import (
	"fmt"
	"go-quizz/m/internal/transport/httpx/dto"
	"net/http"
)

func (h *Handler) GetAllLobbies(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.GetAllLobbiesResponse]()

	lobbies := h.Lobby.GetAllLobbies()

	// If there are no lobbies, return a 404
	if len(lobbies) == 0 {
		apiResponse.Code = http.StatusNotFound
		apiResponse.Message = "no lobbies found"

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	getAllLobbiesResponse := dto.GetAllLobbiesResponse{}

	// creating the response for every found lobby
	for _, lobby := range lobbies {
		getAllLobbiesResponse.Lobbies = append(getAllLobbiesResponse.Lobbies, dto.GetLobbyResponse{
			ID:       lobby.ID,
			HostID:   lobby.HostID,
			Clients:  lobby.Clients,
			Messages: lobby.Messages,
		})
	}

	apiResponse.Data = getAllLobbiesResponse

	// sending the response
	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}

func (h *Handler) GetLobby(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.GetLobbyResponse]()

	// lobby not found because ID is invalid
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("get lobby[%v]: %w", lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// lobby not found
	lobby, err := h.Lobby.GetLobby(lobbyID)
	if err != nil {
		apiResponse.Code = http.StatusNotFound
		apiResponse.Message = fmt.Errorf("get lobby[%v]: %w", lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// lobby found
	apiResponse.Data = dto.GetLobbyResponse{
		ID:       lobby.ID,
		HostID:   lobby.HostID,
		Clients:  lobby.Clients,
		Messages: lobby.Messages,
	}

	// sending the response
	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}

func (h *Handler) PostLobby(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.CreateLobbyResponse]()

	// lobby already exists
	lobby_id, err := h.Lobby.OpenLobby()
	if err != nil {
		apiResponse.Code = http.StatusConflict
		apiResponse.Message = fmt.Errorf("opening lobby: %w", err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	apiResponse.Data = dto.CreateLobbyResponse{
		ID: lobby_id,
	}

	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteLobby(writer http.ResponseWriter, req *http.Request) {
	apiResponse := dto.NewAPIResponse[dto.DeleteLobbyResponse]()

	// uuid is invalid
	lobbyID, err := getUUIDFromUri(req, "lobby_id")
	if err != nil {
		apiResponse.Code = http.StatusBadRequest
		apiResponse.Message = fmt.Errorf("closing lobby[%v]: %w", lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// lobby doesn't exist
	if err := h.Lobby.CloseLobby(lobbyID); err != nil {
		apiResponse.Code = http.StatusNotFound
		apiResponse.Message = fmt.Errorf("closing lobby[%v]: %w", lobbyID, err).Error()

		if err := encodeResponse(writer, apiResponse); err != nil {
			panic(err)
		}

		return
	}

	// no data

	if err := encodeResponse(writer, apiResponse); err != nil {
		panic(err)
	}
}
