package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"go-quizz/m/frontend"
	"go-quizz/m/internal/transport/web/handlers"
)

func main() {
	handler := handlers.NewHandler()

	staticSub, err := fs.Sub(frontend.Static, "static")
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.FS(staticSub))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("POST /lobby", handler.LobbyPage)

	fmt.Println("Listening on localhost:8443")
	http.ListenAndServe(":8443", nil)
}
