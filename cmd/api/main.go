package main

import (
	"fmt"
	"net/http"

	"go-quizz/m/internal/core/services/client"
	"go-quizz/m/internal/core/services/lobby"
	"go-quizz/m/internal/transport/httpx"
)

func main() {
	lobbySrvc := lobby.NewService()
	clientSrvc := client.NewService()

	handler := httpx.NewHandler(lobbySrvc, clientSrvc)

	fmt.Println("Listening on localhost:8080...")
	http.ListenAndServe(":8080", handler.Router)

	// https://github.com/TutorialEdge/go-rest-api-course
}
