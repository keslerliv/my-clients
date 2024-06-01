package handlers

import (
	"html/template"
	"net/http"
)

// Client GET handler
func HomeGet(w http.ResponseWriter, r *http.Request) {

	// parse template
	tmpl, err := template.ParseFiles("internal/handlers/templates/index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// execute template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
