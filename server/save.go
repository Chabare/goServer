package server

import "net/http"

// SaveHandler TODO: Comment
func SaveHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/save/"):]

	body := req.PostFormValue("body")
	page := Page{title, []byte(body)}
	page.Save()

	http.Redirect(w, req, "/", http.StatusFound)
}
