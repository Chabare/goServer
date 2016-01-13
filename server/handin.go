package server

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

// HandinHandler TODO: Comment
func HandinHandler(w http.ResponseWriter, r *http.Request) {
	names := []string{"Torben", "Patrick"}

	_, weekNr := time.Now().ISOWeek()

	t, err := template.ParseFiles(filepath.Join("resources/", "templates/", "handin.html"))

	if err != nil {
		log.Println(err)
	}

	if time.Now().Weekday() > time.Tuesday || time.Now().Weekday() == time.Sunday {
		weekNr++
	}
	turn := weekNr % 2

	t.Execute(w, struct{ Name string }{names[turn]})
}
