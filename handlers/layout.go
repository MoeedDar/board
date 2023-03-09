package handlers

import (
	"board/globals"
	"html/template"
	"net/http"
	"path/filepath"
)

func Layout(w http.ResponseWriter, r *http.Request, childPath string, data interface{}) {
	tmpl, err := template.ParseFiles(filepath.Join(globals.TemplatesDir, childPath), filepath.Join(globals.TemplatesDir, "layout.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
