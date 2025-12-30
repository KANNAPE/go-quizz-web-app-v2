package dto

import "github.com/google/uuid"

// Requests
type ClientConnectionRequest struct {
	Username string `validate:"required"`
}

// Responses
type GetClientResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type GetLobbyClientsResponse struct {
	LobbyID uuid.UUID           `json:"lobby_id"`
	Clients []GetClientResponse `json:"connected_clients"`
}

type LobbyClientConnectsResponse struct {
	LobbyID  uuid.UUID `json:"lobby_id"`
	ClientID uuid.UUID `json:"client_id"`
	Username string    `json:"username"`
}

type LobbyClientDisconnectsResponse struct {
}
