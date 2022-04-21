package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"testing"
)

func TestElementsByTagName(t *testing.T) {
	url := "https://golang.org"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	// TODO: printf *html.Node
	fmt.Println("images: ", images)
	fmt.Println("headings: ", headings)
}
