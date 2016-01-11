package server

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, page *Page) {
	t, err := template.ParseFiles(filepath.Join("server/", "templates/", tmpl) + ".html")

	if err != nil {
		log.Println(err)
	}

	t.Execute(w, page)
}
