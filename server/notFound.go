package server

import (
	"io"
	"net/http"
)

// ErrorHandler TODO: DO
func ErrorHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This site doesn't exist.")
}
