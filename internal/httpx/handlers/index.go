package handlers

import (
	"go-quizz/m/frontend"
	"net/http"
	"text/template"
)

func HomePage(writer http.ResponseWriter, request *http.Request) {
	homePageTemplate, err := template.ParseFS(frontend.Templates, "templates/layout.html", "templates/index.html")
	if err != nil {
		http.Error(writer, "template error", http.StatusInternalServerError)
		return
	}

	err = homePageTemplate.Execute(writer, nil)
	if err != nil {
		http.Error(writer, "render error", http.StatusInternalServerError)
	}
}
