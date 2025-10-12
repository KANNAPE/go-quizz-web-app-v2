package main

import (
	"fmt"
	"net/http"

	service "go-quizz/m/internal/core/service/lobby"
	"go-quizz/m/internal/transport/httpx"
)

func main() {
	lobbySrvc := service.NewLobbyService()

	handler := httpx.NewHandler(lobbySrvc)

	fmt.Println("Listening on localhost:8080...")
	http.ListenAndServe(":8080", handler.Router)

	// https://github.com/TutorialEdge/go-rest-api-course
}
