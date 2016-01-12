package server

import (
	"net"
	"net/http"
	"path/filepath"
	"text/template"
)

// IndexPage TODO: Comment
type IndexPage struct {
	Title    string
	CSSFiles []string
	Header   []Header
	Location Location
}

// IndexHandler TODO: Comment
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
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
	}
	t := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles(filepath.Join("resources/", "templates/", "index")+".html", "resources/templates/weather.tpl"))

	loc := getLocation(net.ParseIP(r.RemoteAddr))

	page := IndexPage{"Index", []string{"main", "header", "links"}, ReadLinkList(), loc}

	t.Execute(w, page)
}
