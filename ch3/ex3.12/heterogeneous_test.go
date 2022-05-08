package heterogeneous

import "testing"

func TestHeterogeneous(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		want   bool
	}{
		{
			"abcd",
			"acbd",
			true,
		},
		{
			"aabcd",
			"accbd",
			false,
		},
		{
			"",
			"",
			true,
		},
	}

	for _, test := range tests {
		got := check(test.input1, test.input2)
		if got != test.want {
			// t.Errorf("check(%s, %s) = %s, expect %s\n", test.input1, test.input2, btoa(got), btoa(test.want))
			t.Errorf("check(%s, %s) = %t, expect %t\n", test.input1, test.input2, got, test.want)
		}
	}
}

func btoa(val bool) string {
	if val {
		return "true"
	}
	return "false"
}
