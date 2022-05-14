package unique

import "testing"

func TestUnique(t *testing.T) {
	var tests = []struct {
		input []string
		want  []string
	}{
		{
			[]string{"one", "two", "two", "three"},
			[]string{"one", "two", "three"},
		},
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"one", "two", "two", "three", "4", "4", "5"},
			[]string{"one", "two", "three", "4", "5"},
		},
	}

	sliceEqual := func(lhs, rhs []string) bool {
		len1, len2 := len(lhs), len(rhs)
		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if lhs[i] != rhs[i] {
				return false
			}
		}
		return true
	}

	for _, test := range tests {
		got := make([]string, len(test.input))
		copy(got, test.input)
		got = unique(got)
		if !sliceEqual(got, test.want) {
			t.Errorf("unique(%v) = %v,  expect %v", test.input, got, test.want)
		}
	}
}
