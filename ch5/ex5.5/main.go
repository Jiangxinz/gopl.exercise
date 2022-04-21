package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages: %v\n", err)
			continue
		}
		fmt.Fprintf(os.Stdout, "words: %d, images: %d\n", words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	var q []*html.Node
	q = append(q, n)
	for len(q) != 0 {
		n = q[0]
		q = q[1:]

		curWords, curImages := count(n)
		words += curWords
		images += curImages

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			q = append(q, c)
		}
	}
	return
}

func count(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		scan := bufio.NewScanner(strings.NewReader(n.Data))
		scan.Split(bufio.ScanWords)
		for scan.Scan() {
			words++
		}
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	return
}
