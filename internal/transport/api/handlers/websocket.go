package handlers

import (
	"log"
	"net/http"

	"go-quizz/m/internal/transport/websocket"

	"github.com/google/uuid"
)

func (h *Handler) ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	// 1. Parse IDs from URL
	lobbyID, err := getUUIDFromUri(r, "lobby_id")
	if err != nil {
		http.Error(w, "Invalid lobby ID", http.StatusBadRequest)
		return
	}

	clientIDStr := r.URL.Query().Get("client_id")
	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	// 2. Validate that the client actually belongs to this lobby
	if _, err := h.Lobby.GetClientInLobby(lobbyID, clientID); err != nil {
		http.Error(w, "Client not authorized in this lobby", http.StatusForbidden)
		return
	}

	// 3. Upgrade the HTTP connection to a WebSocket
	// Note: We need to export 'upgrader' from your websocket package or define one here.
	// For simplicity, let's assume you make the Upgrader public in client.go (change 'var upgrader' to 'var Upgrader')
	// OR, define a local one here.
	conn, err := websocket.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	// 4. Get the Hub for this lobby
	hub := h.HubMgr.GetHub(lobbyID)

	// 5. Create the Client instance
	client := &websocket.Client{
		ID:   clientID,
		Hub:  hub,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	// 6. Register and start pumps
	client.Hub.Register <- client

	// Allow collection of memory by starting goroutines
	go client.WritePump()
	go client.ReadPump()
}
