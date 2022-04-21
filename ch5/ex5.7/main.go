package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func parse(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("parse: %s, err: %v", url, err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parse: %s, err: %v", url, err)
	}
	helper(doc)
	return nil
}

func helper(n *html.Node) {
	forEachNode(n, start, end)
}

// TODO: enclose into class

var depth int

type HandlerMap map[html.NodeType]func(*html.Node)

var startHandlerMap = HandlerMap{
	html.ElementNode: startElement,
	html.TextNode:    startText,
	html.CommentNode: startComment,
}

var endHandlerMap = HandlerMap{
	html.ElementNode: endElement,
	html.TextNode:    endText,
	html.CommentNode: endComment,
}

func start(n *html.Node) {
	run(n, startHandlerMap)
	depth++
}

func end(n *html.Node) {
	depth--
	run(n, endHandlerMap)
}

func run(n *html.Node, handlerMap HandlerMap) {
	handler, ok := handlerMap[n.Type]
	if !ok {
		return
	}
	handler(n)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	attrs := make([]string, 0, len(n.Attr))
	attrs = append(attrs, "")
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s:"%s"`, a.Key, a.Val))
	}

	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}
	fmt.Fprintf(os.Stdout, "%*s<%s%s%s\n", depth*2, "", n.Data, strings.Join(attrs, " "), end)
}

func endElement(n *html.Node) {
	if n.FirstChild == nil {
		return
	}
	fmt.Fprintf(os.Stdout, "%*s</%s>\n", depth*2, "", n.Data)
}

func startText(n *html.Node) {
	if len(strings.TrimSpace(n.Data)) == 0 {
		return
	}
	fmt.Fprintf(os.Stdout, "%*s%s\n", depth*2, "", n.Data)
}

func endText(n *html.Node) {
	endElement(n)
}

func startComment(n *html.Node) {
	fmt.Fprintf(os.Stdout, "<!--%s-->\n", n.Data)
}

func endComment(n *html.Node) {
	endElement(n)
}

// func main() {
// 	for _, url := range os.Args[1:] {
// 		parse(url)
// 	}
// }
