package dto

import "go-quizz/m/internal/core/domain"

type Client struct {
	ID       string `json:"client_id"`
	Username string `json:"username"`
}

type Lobby struct {
	ID      string   `json:"lobby_id"`
	Clients []Client `json:"clients"`
	HostID  string   `json:"host_id"`
	Chat    ChatRoom `json:"chat_room"`
}

func (clientDto *Client) FromClient(client domain.Client) {

}

func (lobbyDto *Lobby) FromLobby(lobby domain.Lobby) {

}
