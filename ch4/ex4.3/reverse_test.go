package reverse

import "testing"

func TestReverse(t *testing.T) {
	var tests = []struct {
		input *[ARRLEN]int
		want  *[ARRLEN]int
	}{
		{
			&[ARRLEN]int{1, 3, 5, 7, 9},
			&[ARRLEN]int{9, 7, 5, 3, 1},
		},
	}

	for _, test := range tests {
		got := &[ARRLEN]int{}
		copy(got[:], test.input[:])
		reverse(got)
		if *got != *test.want {
			t.Errorf("reverse(%v) = %v, expect %v\n", test.input, got, test.want)
		}
	}
}
