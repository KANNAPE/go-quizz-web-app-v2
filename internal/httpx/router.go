package httpx

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	//handler := transport.NewHandler()

	// routes
	router.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, "API Endpoint Hit")
	})
	// router.HandleFunc("/api/lobby", ).Methods("POST")
	// router.HandleFunc("/api/lobby/{id}", handler.LobbyPage)

	return router
}
