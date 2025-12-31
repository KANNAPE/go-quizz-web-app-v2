package handlers

import "net/http"

func (h *Handler) LobbyPage(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	lobbyIdCookie, err := request.Cookie("lobby_id")

	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"Username": username,
		"LobbyID":  lobbyIdCookie.Value,
	}

	h.Templates["lobby.html"].Execute(writer, data)

	// deleting the cookie
	lobbyIdCookie.MaxAge = -1
}
