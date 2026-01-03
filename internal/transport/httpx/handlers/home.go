package handlers

import "net/http"

func (h *Handler) CreateLobbyPage(writer http.ResponseWriter, request *http.Request) {
	h.Templates["index.html"].Execute(writer, nil)
}

func (h *Handler) JoinLobbyPage(writer http.ResponseWriter, request *http.Request) {
	lobbyId := request.URL.Query().Get("id")
	if lobbyId == "" {
		// redirect to index.html
		http.Redirect(writer, request, "/", http.StatusSeeOther)

		return
	}

	data := map[string]interface{}{
		"JoinLobby": true,
	}

	h.Templates["index.html"].Execute(writer, data)
}
