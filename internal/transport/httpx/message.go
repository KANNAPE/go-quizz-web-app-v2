package httpx

import (
	"net/http"
)

func (h *Handler) GetLobbyMessages(writer http.ResponseWriter, req *http.Request) {
	// messages := h.Lobby.()

	// if err := json.NewEncoder(writer).Encode(messages); err != nil {
	// 	panic(err)
	// }
}

func (h *Handler) GetMessage(writer http.ResponseWriter, req *http.Request) {
	// messageID, err := getUUIDFromUri(req, "id")
	// if err != nil {
	// 	writer.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// message, err := h.Message.Get(messageID)
	// if err != nil {
	// 	writer.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// if err := json.NewEncoder(writer).Encode(message); err != nil {
	// 	panic(err)
	// }
}

type PostMessageRequest struct {
	Body string `json:"content" validate:"required"`
}

func (h *Handler) PostMessage(writer http.ResponseWriter, req *http.Request) {
	// var messageReq PostMessageRequest
	// if err := json.NewDecoder(req.Body).Decode(&messageReq); err != nil {
	// 	writer.WriteHeader(http.StatusUnprocessableEntity)
	// 	return
	// }

	// message_id, err := h.Message.Create(messageReq.Body)
	// if err != nil {
	// 	writer.WriteHeader(http.StatusUnprocessableEntity)
	// 	return
	// }

	// if err := json.NewEncoder(writer).Encode(message_id); err != nil {
	// 	panic(err)
	// }
}
