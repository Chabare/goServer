package server

import "net/http"

// EditHandler TODO: Comment
func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	page, err := LoadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, "edit", page)
}
