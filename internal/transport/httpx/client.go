package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) GetClients(writer http.ResponseWriter, req *http.Request) {
	clients := h.Client.GetAll()

	if len(clients) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(writer).Encode(clients); err != nil {
		panic(err)
	}
}

func (h *Handler) GetClient(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	client_id := vars["id"]
	if client_id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	clientID, err := uuid.Parse(client_id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	client, err := h.Client.Get(clientID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(writer).Encode(client); err != nil {
		panic(err)
	}
}

type PostClientRequest struct {
	Username string `json:"username" validate:"required"`
}

func (h *Handler) PostClient(writer http.ResponseWriter, req *http.Request) {
	var clientReq PostClientRequest
	if err := json.NewDecoder(req.Body).Decode(&clientReq); err != nil {
		writer.WriteHeader(http.StatusNotAcceptable)
		return
	}

	validate := validator.New()
	err := validate.Struct(clientReq)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	client_id := h.Client.Register(clientReq.Username)

	if err := json.NewEncoder(writer).Encode(client_id); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
