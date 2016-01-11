package server

import (
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Link TODO: Comment
type Link struct {
	Name string
	URL  *url.URL
}

// ByLinkName TODO: Comment
type ByLinkName []Link

// ByHeaderName TODO: Comment
type ByHeaderName []Header

func (links ByLinkName) Len() int      { return len(links) }
func (links ByLinkName) Swap(i, j int) { links[i], links[j] = links[j], links[i] }
func (links ByLinkName) Less(i, j int) bool {
	return (strings.Compare(links[i].Name, links[j].Name) < 0)
}
func (links ByHeaderName) Len() int      { return len(links) }
func (links ByHeaderName) Swap(i, j int) { links[i], links[j] = links[j], links[i] }
func (links ByHeaderName) Less(i, j int) bool {
	return (strings.Compare(links[i].Name, links[j].Name) < 0)
}

// Header TODO: Comment
type Header struct {
	Name  string
	Links []Link
}

// ReadLinkList TODO: Comment
func ReadLinkList() []Header {
	header := make([]Header, 9, 9)

	lines := getLines("server/links")

	headerCount := -1
	var linkCount int
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			headerCount++
			header[headerCount].Links = make([]Link, 3, 3)
			header[headerCount].Name = strings.TrimSpace(strings.TrimLeft(line, "#"))
			linkCount = 0
		} else if len(strings.TrimSpace(line)) > 0 {
			linkName, url := getLinkNameAndURL(line)
			header[headerCount].Links[linkCount] = Link{linkName, url}
			linkCount++
		}
	}

	return sortHeader(header)
}

func sortHeader(header []Header) []Header {
	hn := make(ByHeaderName, 9, 9)
	for i, l := range header {
		hn[i] = l
	}
	sorted := sortHeaderNames(hn)
	for i := range sorted {
		header[i] = sorted[i]
	}

	header = trans(header)

	// Sort links in each header
	for _, h := range header {
		ln := make(ByLinkName, 3, 3)
		for i, l := range h.Links {
			ln[i] = l
		}
		sorted := sortLinkNames(ln)
		for i := range sorted {
			h.Links[i] = sorted[i]
		}
	}

	return header
}

func trans(header []Header) []Header {
	for ind := 0; ind < 3; ind++ {
		header[ind], header[ind*3] = header[ind*3], header[ind]
	}

	return header
}

func sortLinkNames(names ByLinkName) ByLinkName {
	sort.Sort(names)

	return names
}

func sortHeaderNames(names ByHeaderName) ByHeaderName {
	sort.Sort(names)

	return names
}

func getLinkNameAndURL(line string) (string, *url.URL) {
	name := line[:strings.Index(line, ":")]
	url, _ := url.Parse(line[strings.Index(line, ":")+1:])

	return name, url
}

func getLines(filename string) []string {
	file := open(filename)
	defer file.Close()

	data := make([]byte, 4096)
	file.Read(data)
	lines := strings.Split(string(data), "\n")

	// Check for empty byte array
	i, _ := strconv.Atoi(lines[len(lines)-1])
	if i == 0 {
		lines = lines[:len(lines)-1]
	}

	return lines[:]
}

// open opens the file
func open(name string) *os.File {
	f, err := os.Open(name)

	if err != nil {
		log.Fatal(err)
	}

	return f
}
