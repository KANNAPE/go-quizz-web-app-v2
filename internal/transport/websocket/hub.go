package websocket

import (
	"encoding/json"
	"log"

	"go-quizz/m/internal/core/domain"

	"github.com/google/uuid"
)

type MessageService interface {
	CreateMessage(lobbyID uuid.UUID, senderID uuid.UUID, body string) (uuid.UUID, error)
	GetLobbyMessage(lobbyID uuid.UUID, messageID uuid.UUID) (domain.Message, error)
}

type IncomingMessage struct {
	Client *Client
	Body   []byte
}

type Hub struct {
	LobbyID    uuid.UUID
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Inbox      chan IncomingMessage
	Service    MessageService
}

func NewHub(lobbyID uuid.UUID, service MessageService) *Hub {
	return &Hub{
		LobbyID:    lobbyID,
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Inbox:      make(chan IncomingMessage),
		Service:    service,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case msg := <-h.Inbox:
			// parse Request
			var req struct {
				Body string `json:"body"`
			}
			if err := json.Unmarshal(msg.Body, &req); err != nil {
				continue
			}

			// call Service
			msgID, err := h.Service.CreateMessage(h.LobbyID, msg.Client.ID, req.Body)
			if err != nil {
				log.Printf("Service error: %v", err)
				continue
			}

			// get message
			fullMsg, _ := h.Service.GetLobbyMessage(h.LobbyID, msgID)

			// broadcast
			bytes, _ := json.Marshal(fullMsg)
			for client := range h.Clients {
				select {
				case client.Send <- bytes:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
