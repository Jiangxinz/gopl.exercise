package mystringreader

import (
	"fmt"
	"golang.org/x/net/html"
	"testing"
)

func TestMyStringReader(t *testing.T) {
	var s string = "<html><body><p>hi</p></body></html>"
	reader, err := MyStringReader(s)
	if err != nil {
		t.Errorf("MyStringReader: %v\n", err)
	}
	doc, err := html.Parse(reader)
	if err != nil {
		t.Errorf("html.Parse: %v\n", err)
	}
	// fmt.Println(doc)
	visit(doc)
}

func visit(n *html.Node) {
	if n != nil {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}

func TestSliceParameter(t *testing.T) {
	var s string = "<html><body><p>hi</p></body></html>"
	reader, err := MyStringReader(s)
	if err != nil {
		t.Errorf("MyStringReader: %v\n", err)
	}
	var buf []byte = make([]byte, 1024, 1024)
	n, err := reader.Read(buf)
	fmt.Println("buf:", buf[:n])
}
