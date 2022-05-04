package eval

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	exprs := []string{
		"sin(-x)*pow(1.5, -r)",
		"pow(2, sin(y))*pow(2, sin(x))/12",
		"sin(-x)*pow(1.5, -r)*pow(2, sin(x))",
		"sin(-x)",
		"-x",
		"sin(1+x)",
	}
	for _, expr := range exprs {
		e, _ := Parse(expr)
		fmt.Println(e)
	}
}
