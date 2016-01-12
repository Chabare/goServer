package server

import (
	"net/url"
	"strings"
)

// Link ...
type Link struct {
	Name string
	URL  *url.URL
}

// Header ...
type Header struct {
	Name  string
	Links []Link
}

// Links ...
type Links []Link

// Headers ...
type Headers []Header

// Len ...
func (l Links) Len() int { return len(l) }

// Len ...
func (h Headers) Len() int { return len(h) }

// Swap ...
func (l Links) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

// Swap ...
func (h Headers) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// ByLink ...
type ByLink struct{ Links }

// ByHeader ...
type ByHeader struct{ Headers }

// Less ...
func (l Links) Less(i, j int) bool { return strings.Compare(l[i].Name, l[j].Name) < 0 }

// Less ...
func (h Headers) Less(i, j int) bool { return strings.Compare(h[i].Name, h[j].Name) < 0 }
