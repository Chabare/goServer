package server

import "net/http"

// ViewHandler TODO: Comment
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	page, err := LoadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	// Replace new lines with break tag
	// page.Body = []byte(strings.Replace(string(page.Body), string([]byte{13}), "<br>", 1024))
	renderTemplate(w, "view", page)
}
