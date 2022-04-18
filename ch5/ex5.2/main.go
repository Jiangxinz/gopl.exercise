package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func count(result map[string]uint32, n *html.Node) {
	if n.Type == html.ElementNode {
		result[n.Data] += 1
	}
}

func helper(result map[string]uint32, n *html.Node) {
	var q []*html.Node
	q = append(q, n)
	for len(q) != 0 {
		n = q[0]
		q = q[1:]
		count(result, n)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			q = append(q, c)
		}
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse: %v\n", err)
		os.Exit(1)
	}
	var result map[string]uint32 = make(map[string]uint32)
	for d := doc; d != nil; d = d.NextSibling {
		helper(result, d)
	}
	fmt.Println(result)
}
