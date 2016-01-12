package server

import (
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ReadLinkList TODO: Comment
func ReadLinkList() []Header {
	header := make([]Header, 9, 9)

	lines := getLines("resources/links")

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
	sortHeaderNames(header)

	header = trans(header)

	// Sort links in each header
	for _, h := range header {
		sortLinkNames(h.Links)
	}

	return header
}

func trans(header []Header) []Header {
	for ind := 0; ind < 3; ind++ {
		header[ind], header[ind*3] = header[ind*3], header[ind]
	}

	return header
}

func sortLinkNames(names []Link) {
	sort.Sort(ByLink{names})
}

func sortHeaderNames(names []Header) {
	sort.Sort(ByHeader{names})
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
