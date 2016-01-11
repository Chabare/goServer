package server

import "io/ioutil"

// Page represents a webpage with its title and body
type Page struct {
	Title string
	Body  []byte
}

// Save saves the content of the page to a textfile
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage loads the content of the page from a textfile
func LoadPage(title string) (*Page, error) {
	body, err := ioutil.ReadFile(title)

	return &Page{title, body}, err
}
