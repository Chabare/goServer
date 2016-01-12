package server

import (
	"html"
	"io/ioutil"
	"strings"
)

// Page represents a webpage with its title and body
type Page struct {
	Title string
	Body  []byte
}

// Save saves the content of the page to a textfile
func (p *Page) Save() error {
	filename := "resources/" + p.Title
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage loads the content of the page from a textfile
func LoadPage(title string) (*Page, error) {
	filename := title
	var body []byte
	var err error
	if !strings.Contains(filename, "/") {
		filename = "resources/" + filename
	}
	if strings.HasPrefix(filename, "templates") {
		filename = "resources/" + filename
		body, err = ioutil.ReadFile(filename)
		body = []byte(html.EscapeString(string(body)))
	} else {
		body, err = ioutil.ReadFile(filename)
	}

	return &Page{title, body}, err
}
