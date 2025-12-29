package main

import (
	"fmt"
	"go-quizz/m/frontend"
	"html/template"
	"io/fs"
	"net/http"
)

func main() {
	templates := parseTemplates()

	staticSub, err := fs.Sub(frontend.Static, "static")
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.FS(staticSub))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates["index.html"].Execute(w, nil)
	})
	http.HandleFunc("POST /lobby", func(w http.ResponseWriter, r *http.Request) {
		templates["lobby.html"].Execute(w, nil)
	})

	fmt.Println("Listening on localhost:8443")
	http.ListenAndServe(":8443", nil)
}

func parseTemplates() map[string]*template.Template {
	cache := map[string]*template.Template{}

	// reading through all files in templates folder
	entries, err := frontend.Templates.ReadDir("templates")
	if err != nil {
		panic(err)
	}

	// parsing templates, using the layout.html as base template
	for _, entry := range entries {
		name := entry.Name()
		if name == "layout.html" {
			continue
		}

		ts, err := template.ParseFS(frontend.Templates, "templates/layout.html", "templates/"+name)
		if err != nil {
			panic(err)
		}

		cache[name] = ts
	}

	return cache
}
