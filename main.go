package main

import (
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/chabare/goServer/server"
)

func main() {
	http.HandleFunc("/css/", server.ViewHandler)
	http.HandleFunc("/img/", server.ViewHandler)
	serveSingle("/favicon.ico", "./favicon.ico")
	http.HandleFunc("/", server.IndexHandler)
	http.HandleFunc("/edit/", server.EditHandler)
	http.HandleFunc("/save/", server.SaveHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func handler(w http.ResponseWriter, req *http.Request) {
	page, err := server.LoadPage("index")

	if err != nil || !regexp.MustCompile("/$").MatchString(req.URL.Path) {
		http.Redirect(w, req, "/", http.StatusFound)
	}

	io.WriteString(w, string(page.Body))
}
