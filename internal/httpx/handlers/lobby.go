package handlers

import (
	"fmt"
	"go-quizz/m/frontend"
	"go-quizz/m/internal/core/services/lobby"
	"net/http"
	"text/template"
)

type test struct {
	JoinURL  string
	LobbyID  string
	Username string
}

var Test test

func (handler *Handler) CreateOrJoinLobbyPage(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		http.Error(writer, "form parsing failed", http.StatusBadRequest)
		return
	}

	username := request.PostFormValue("username")

	newLobby := lobby.GenerateNewLobby(username)
	if newLobby == nil {
		http.Error(writer, "lobby generation error", http.StatusBadRequest)
		return
	}

	lobbyUrl := "/lobby/" + newLobby.ID.String()

	scheme := "http"
	if request.TLS != nil {
		scheme = "https"
	}

	handler.JoinURL = fmt.Sprintf("%v://%v%v/%v", scheme, request.Host, request.RequestURI, newLobby.ID)
	handler.Lobby.FromLobby(newLobby)
	handler.Host.FromClient(newLobby.Clients[newLobby.HostID])

	http.Redirect(writer, request, lobbyUrl, http.StatusSeeOther)
}

func (handler *Handler) LobbyPage(writer http.ResponseWriter, request *http.Request) {
	lobbyPageTemplate, err := template.ParseFS(frontend.Templates, "templates/layout.html", "templates/lobby.html")
	if err != nil {
		http.Error(writer, "template error", http.StatusInternalServerError)
		return
	}

	err = lobbyPageTemplate.Execute(writer, handler)
	if err != nil {
		http.Error(writer, "render error", http.StatusInternalServerError)
	}
}
