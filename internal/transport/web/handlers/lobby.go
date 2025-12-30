package handlers

import "net/http"

func (h *Handler) LobbyPage(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")

	data := map[string]interface{}{
		"Username": username,
		"LobbyID":  "134123123",
	}

	h.Templates["lobby.html"].Execute(writer, data)
}
