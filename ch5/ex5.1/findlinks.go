package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func visitRecursively(links []string, n *html.Node) []string {
	links = appendNode(links, n)

	if n.FirstChild != nil {
		links = visitRecursively(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visitRecursively(links, n.NextSibling)
	}
	return links
}

func appendNode(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visitRecursively(nil, doc) {
		fmt.Println(link)
	}
}
