package handlers

import (
	"go-quizz/m/frontend"
	"net/http"
	"text/template"
)

func LobbyPage(writer http.ResponseWriter, request *http.Request) {
	lobbyPageTemplate, err := template.ParseFS(frontend.Templates, "templates/layout.html", "templates/lobby.html")
	if err != nil {
		http.Error(writer, "template error", http.StatusInternalServerError)
		return
	}

	lobbyID := "AJHDHSG"
	joinURL := "https://app.kannape.fr/lobby/" + lobbyID
	username := "empereur du caca"

	data := struct {
		JoinURL  string
		LobbyID  string
		Username string
	}{
		LobbyID:  lobbyID,
		JoinURL:  joinURL,
		Username: username,
	}

	err = lobbyPageTemplate.Execute(writer, data)
	if err != nil {
		http.Error(writer, "render error", http.StatusInternalServerError)
	}
}
