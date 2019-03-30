package html_link_parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

// Link represents a Link (<a href="...") in HTML
type Link struct {
	Href string
	Text string
}

//Parse will take in an HTML document and will return a
//slice of links parsed from it
func Parse(reader io.Reader) ([]Link, error) {
	doc, err := html.Parse(reader)

	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
		//fmt.Println(node)
	}
	//dfs(doc,"")

	return links, nil
}

func buildLink(node *html.Node) Link {
	var result Link

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			result.Href = attr.Val
			break
		}
	}

	result.Text = text(node)

	return result
}

func text(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	if node.Type != html.ElementNode {
		return ""
	}
	var result string

	for child := node.FirstChild; child != nil; child = child.NextSibling {

		result += text(child)
	}

	return strings.Join(strings.Fields(result), " ")

}

func linkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var results []*html.Node

	for child := node.FirstChild; child != nil; child = child.NextSibling {

		results = append(results, linkNodes(child)...)
	}

	return results
}

func dfs(node *html.Node, padding string) {
	msg := node.Data
	if node.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		dfs(child, padding+" ")
	}

}
