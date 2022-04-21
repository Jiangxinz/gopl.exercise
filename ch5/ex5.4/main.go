package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

// 图片: n.Data == "img", a.Key == "src"
// 脚本: n.Data == "script", a.Key == "src"
// 样式表: n.Data == "link", a.Key == "href"

func helper(links []string, n *html.Node, nodeData, attrKey string) []string {
	if n.Type == html.ElementNode && n.Data == nodeData {
		for _, a := range n.Attr {
			if a.Key == attrKey {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = helper(links, c, nodeData, attrKey)
	}
	return links
}

func visit(links []string, n *html.Node, nodeData, attrKey string) {
	for _, link := range helper(nil, n, nodeData, attrKey) {
		fmt.Println(link)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("------------link------------")
	visit(nil, doc, "a", "href")
	fmt.Println("------------img------------")
	visit(nil, doc, "img", "src")
	fmt.Println("------------script------------")
	visit(nil, doc, "script", "src")
	fmt.Println("------------style sheet------------")
	visit(nil, doc, "link", "href")
}
