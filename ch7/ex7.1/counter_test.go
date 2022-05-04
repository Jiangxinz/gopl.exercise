package counter

import (
	"fmt"
	"testing"
)

func TestScanWords(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{
			"This is a Test",
			4,
		},
		{
			"This is a Test\r\n",
			4,
		},
		{
			"This is a Test \r\n",
			4,
		},
	}

	for _, test := range tests {
		var cnt WordCounter
		fmt.Fprintf(&cnt, test.input)
		if int(cnt) != test.want {
			t.Errorf("WordCounter(%s) = %d, expect %d\n", test.input, int(cnt), test.want)
		}
	}
}

func TestScanLines(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{
			"This is a Test",
			1,
		},
		{
			"This is a Test\r\n",
			1,
		},
		{
			"This is a Test \r\nAnother Line",
			2,
		},
		{
			"\r\n",
			1,
		},
	}

	for _, test := range tests {
		var cnt LineCounter
		fmt.Fprintf(&cnt, test.input)
		if int(cnt) != test.want {
			t.Errorf("WordCounter(%s) = %d, expect %d\n", test.input, int(cnt), test.want)
		}
	}
}
