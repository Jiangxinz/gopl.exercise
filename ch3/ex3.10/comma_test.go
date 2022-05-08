package comma

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{
			"123",
			"123",
		},
		{
			"1",
			"1",
		},
		{
			"1123",
			"1,123",
		},
		{
			"",
			"",
		},
		{
			"123a123",
			"123a123",
		},
		{
			"123123",
			"123,123",
		},
	}

	for _, test := range tests {
		got, _ := Comma(test.input)
		if got != test.want {
			t.Errorf("Comma(%s) = %s, expect %s\n", test.input, got, test.want)
		}
	}
}
