package main

import (
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var q, ret []*html.Node
	// fmt.Println("doc.NextSibling: ", doc.NextSibling)
	q = append(q, doc)
	for len(q) != 0 {
		cq := q
		q = make([]*html.Node, 0)
		for _, n := range cq {
			for _, keyWord := range name {
				if n.Data == keyWord {
					ret = append(ret, n)
					break
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				q = append(q, c)
			}
		}
	}
	return ret
}
