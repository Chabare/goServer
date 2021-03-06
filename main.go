package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/chabare/goServer/server"
)

func main() {
	// Single files
	serveSingle("/favicon.ico", "./favicon.ico")
	// Handlers
	http.HandleFunc("/css/", server.ViewHandler)
	http.HandleFunc("/img/", server.ViewHandler)
	http.HandleFunc("/resources/", server.ViewHandler)
	// http.HandleFunc("/auth/login", server.LoginHandler)
	http.HandleFunc("/abgabe/", server.HandinHandler)
	http.HandleFunc("/", server.IndexHandler)
	http.HandleFunc("/edit/", server.EditHandler)
	http.HandleFunc("/save/", server.SaveHandler)

	fmt.Println("Server is running")

	go http.ListenAndServe(":80", nil)
	// Listen on default http port
	httpsErr := http.ListenAndServeTLS(":443", "ssl/www_chabare_me.crt", "ssl/myserver.key", nil)
	if httpsErr != nil {
		log.Fatal("ListenAndServe: ", httpsErr)
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
