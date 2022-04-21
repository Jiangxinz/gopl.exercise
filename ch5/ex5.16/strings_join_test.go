package main

import "testing"

func TestJoin(t *testing.T) {
	type data struct {
		strs []string
		sep  string
	}
	var tests = []struct {
		input data
		want  string
	}{
		{
			data{
				[]string{"a", "b", "c"},
				"-",
			},
			"a-b-c",
		},
		{
			data{
				[]string{"aaa", " ", "c"},
				"*-*",
			},
			"aaa*-* *-*c",
		},
		{
			data{
				[]string{"aaa", " ", "c"},
				"",
			},
			"aaa c",
		},
		{
			data{
				[]string{"", "", ""},
				"",
			},
			"",
		},
		{
			data{
				[]string{""},
				"",
			},
			"",
		},
		{
			data{
				[]string{"abc"},
				" ",
			},
			"abc",
		},
		{
			data{
				[]string{},
				" ",
			},
			"",
		},
	}

	for _, test := range tests {
		got := Join(test.input.sep, test.input.strs...)
		if got != test.want {
			t.Errorf("Join(%s, %s) = %s, want %s\n", test.input.strs, test.input.sep, got, test.want)
		}
	}
}
