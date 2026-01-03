package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"go-quizz/m/frontend"
	"go-quizz/m/internal/transport/httpx/handlers"
)

func main() {
	handler := handlers.NewHandler()

	staticSub, err := fs.Sub(frontend.Static, "static")
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.FS(staticSub))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler.CreateLobbyPage)
	http.HandleFunc("GET /lobby", handler.JoinLobbyPage)
	http.HandleFunc("POST /lobby", handler.InLobbyPage)
	http.HandleFunc("/lobby/ws", handler.InLobbyWebsocketConnection)

	fmt.Println("Listening on localhost:8443")
	http.ListenAndServe(":8443", nil)
}
