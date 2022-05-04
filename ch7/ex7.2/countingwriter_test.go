package countingwriter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var tests = []struct {
		input string
		want  int64
	}{
		{
			"Computer Architecture",
			int64(len("Computer Architecture")),
		},
		{},
	}
	for _, test := range tests {
		var buf bytes.Buffer
		w, p := CountingWriter(&buf)
		fmt.Fprintf(w, test.input)
		if int(*p) != len(test.input) {
			t.Errorf("*p = %d, expect %d\n", *p, len(test.input))
		}
	}
}
