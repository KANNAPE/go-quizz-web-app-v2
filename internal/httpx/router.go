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

	handler := handlers.NewHandler()

	// routes
	router.HandleFunc("/", handler.HomePage)
	router.HandleFunc("/lobby", handler.CreateOrJoinLobbyPage).Methods("POST")
	router.HandleFunc("/lobby/{id}", handler.LobbyPage)

	return router
}
