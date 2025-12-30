package main

import (
	"fmt"
	"net/http"

	"go-quizz/m/internal/core/services/lobby"
	"go-quizz/m/internal/transport/api/handlers"
)

func main() {
	lobbySrvc := lobby.NewService()

	handler := handlers.NewHandler(lobbySrvc)

	fmt.Println("Listening on localhost:8080...")
	http.ListenAndServe(":8080", handler)

	// https://github.com/TutorialEdge/go-rest-api-course
}
