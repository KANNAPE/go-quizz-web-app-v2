package httpx

import (
	"go-quizz/m/internal/httpx/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	var handler handlers.Handler

	router.HandleFunc("/", handler.HomePage)

	return router
}
