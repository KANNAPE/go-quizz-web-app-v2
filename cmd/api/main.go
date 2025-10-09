package main

import (
	"fmt"
	"go-quizz/m/internal/httpx"
	"net/http"
)

func main() {
	router := httpx.NewRouter()

	fmt.Println("Listening on localhost:8080...")
	http.ListenAndServe(":8080", router)

	// https://github.com/TutorialEdge/go-rest-api-course
}
