package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) InLobbyPage(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")

	data := map[string]interface{}{
		"Username": username,
	}

	h.Templates["lobby.html"].Execute(writer, data)
}

func (h *Handler) InLobbyWebsocketConnection(writer http.ResponseWriter, request *http.Request) {
	// upgrade the connection to a websocket connection is requested
	if websocket.IsWebSocketUpgrade(request) == false {
		return
	}

	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}

	go func() {
		defer conn.Close()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				return
			}

			if err := conn.WriteMessage(messageType, message); err != nil {
				return
			}
		}
	}()
}
