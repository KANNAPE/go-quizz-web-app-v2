package dto

import "go-quizz/m/internal/core/domain"

type Client struct {
	ID       string `json:"client_id"`
	Username string `json:"username"`
}

type Lobby struct {
	ID      string            `json:"lobby_id"`
	Clients map[string]Client `json:"clients"`
	HostID  string            `json:"host_id"`
	Chat    ChatRoom          `json:"chat_room"`
}

func (clientDto *Client) FromClient(client *domain.Client) {
	clientDto.ID = client.ID.String()
	clientDto.Username = client.Username
}

func (lobbyDto *Lobby) FromLobby(lobby *domain.Lobby) {
	lobbyDto.ID = lobby.ID.String()
	lobbyDto.HostID = lobby.HostID.String()

	lobbyDto.Clients = make(map[string]Client)

	// Clients
	for _, client := range lobby.Clients {
		clientDto := Client{}
		clientDto.FromClient(client)

		lobbyDto.Clients[clientDto.ID] = clientDto
	}

	// Chat room
	lobbyDto.Chat = ChatRoom{}
	lobbyDto.Chat.FromChatRoom(lobby.Chat)
}
