package handlers

import "net/http"

func (h *Handler) HomePage(writer http.ResponseWriter, request *http.Request) {
	h.Templates["index.html"].Execute(writer, nil)
}
