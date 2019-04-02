package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/samueldaviddelacruz/golang-exercises/html-link-parser"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlSet struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")

	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	toXml := urlSet{
		Xmlns: xmlns,
	}

	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	xmlEnconder := xml.NewEncoder(os.Stdout)
	fmt.Print(xml.Header)
	xmlEnconder.Indent("", "  ")
	if err := xmlEnconder.Encode(toXml); err != nil {
		panic(err)
	}
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var queue map[string]struct{}
	nextQueue := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		queue, nextQueue = nextQueue, make(map[string]struct{})
		if len(queue) == 0 {
			break
		}
		for url, _ := range queue {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nextQueue[link] = struct{}{}
				}
			}
		}
	}
	result := make([]string, 0, len(seen))
	for url, _ := range seen {
		result = append(result, url)
	}
	return result
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(reader io.Reader, base string) []string {
	links, _ := html_link_parser.Parse(reader)
	var result []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			result = append(result, base+link.Href)

		case strings.HasPrefix(link.Href, "http"):
			result = append(result, link.Href)
		}
	}

	return result
}

func filter(links []string, keepFn func(string) bool) []string {
	var result []string
	for _, link := range links {
		if keepFn(link) {
			result = append(result, link)
		}
	}
	return result
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

/*
	1. GET the webpage
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links with a different domain
	5. find all the pages (BFS)
    6. print out XML
*/
