package server

import (
	"net/http"
	"path/filepath"
	"reflect"
	"text/template"
)

// Columns ...
type Columns struct {
	Header []Header
}

// Row ...
type Row struct {
	Header  []Header
	Columns []Columns
}

// IndexPage TODO: Comment
type IndexPage struct {
	Title    string
	CSSFiles []string
	Header   []Header
	Location Location
}

// IndexHandler TODO: Comment
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("index.html").Funcs(createIndexFunctionMap()).ParseFiles(filepath.Join("resources/", "templates/", "index")+".html", "resources/templates/weather.tpl"))

	loc := getLocation(r.RemoteAddr)

	page := IndexPage{"Index", []string{"main", "header", "links"}, ReadLinkList(), loc}

	t.Execute(w, page)
}

// NOTE: Use a file for function maps?!
// createIndexFunctionMap TODO: Comment
func createIndexFunctionMap() template.FuncMap {
	return template.FuncMap{
		"mod": func(i int, m int) int {
			return i % m
		},
		"inc": func(i int) int {
			return i + 1
		},
		"divExact": func(i, j int) int {
			if i%j != 0 {
				return -1
			}

			return i / j
		},
		"div": func(i, j int) int {
			return i / j
		},
		"len": func(arr interface{}) int {
			return reflect.ValueOf(arr).Len()
		},
	}
}
