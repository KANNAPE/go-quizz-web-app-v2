package main

import (
	"fmt"
	"net/http"

	"go-quizz/m/internal/core/services/lobby"
	"go-quizz/m/internal/transport/api/handlers"
	"go-quizz/m/internal/transport/websocket"
)

func main() {
	lobbySrvc := lobby.NewService()
	hubMgr := websocket.NewHubManager()

	handler := handlers.NewHandler(lobbySrvc, hubMgr)

	fmt.Println("Listening on localhost:8080...")
	http.ListenAndServe(":8080", handler)
}
