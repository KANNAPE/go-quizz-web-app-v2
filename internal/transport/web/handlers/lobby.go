package handlers

import "net/http"

func (h *Handler) LobbyPage(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")

	data := map[string]interface{}{
		"Username": username,
	}

	h.Templates["lobby.html"].Execute(writer, data)
}
