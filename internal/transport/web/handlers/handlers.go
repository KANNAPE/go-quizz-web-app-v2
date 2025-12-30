package handlers

import (
	"go-quizz/m/internal/transport/web"
	"html/template"
	"net/http"
)

type Handler struct {
	Router    *http.ServeMux
	Templates map[string]*template.Template
}

func NewHandler() *Handler {
	return &Handler{
		Router:    http.NewServeMux(),
		Templates: web.ParseTemplates(),
	}
}
