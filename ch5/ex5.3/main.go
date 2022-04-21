package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func visit(lists []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		for _, a := range n.Attr {
			lists = append(lists, a.Val)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		lists = visit(lists, c)
	}
	return lists
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range visit(nil, doc) {
		fmt.Println(list)
	}
}
