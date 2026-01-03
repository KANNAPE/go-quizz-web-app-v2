package httpx

import (
	"go-quizz/m/frontend"
	"html/template"
)

func ParseTemplates() map[string]*template.Template {
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
