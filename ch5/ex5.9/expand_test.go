package main

import (
	"testing"
)

func TestExpand(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"123$foo123", "123aaa123"},
		{"$foo", "aaa"},
		{"1$foo", "1aaa"},
		{"$foo2", "aaa2"},
		{"$foo$foo", "aaaaaa"},
	}

	for _, test := range tests {
		got := expand(test.input, func(s string) string {
			ret := make([]rune, len(s))
			for i, _ := range s {
				ret[i] = 'a'
			}
			return string(ret)
		})

		if got != test.want {
			t.Errorf("expand(%q) = %v", test.input, got)
		}
	}
}
