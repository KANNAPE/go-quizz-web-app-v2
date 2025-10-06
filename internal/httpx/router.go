package httpx

import (
	"go-quizz/m/frontend"
	"go-quizz/m/internal/httpx/handlers"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	staticSub, err := fs.Sub(frontend.Static, "static")
	if err != nil {
		return nil
	}

	fileServer := http.FileServer(http.FS(staticSub))
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", fileServer))

	// routes
	router.HandleFunc("/", handlers.HomePage)
	router.HandleFunc("/lobby", handlers.CreateOrJoinLobbyPage).Methods("POST")
	router.HandleFunc("/lobby/{id}", handlers.LobbyPage)

	return router
}
