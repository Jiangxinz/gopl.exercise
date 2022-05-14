package reverse

import "testing"

// import "fmt"

func TestReverse(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{
			[]byte("哈12345"),
			[]byte("54321哈"),
		},
		{
			[]byte("哈哈\u00a0\u00a0哈\u00a0哈哈\u00a0\u00a0a"),
			[]byte("a\u00a0\u00a0哈哈\u00a0哈\u00a0\u00a0哈哈"),
		},
		{
			[]byte("哈哈\u00a0\u00a0\u00a0哈\u00a0哈哈\u00a0\u00a0\u00a0a"),
			[]byte("a\u00a0\u00a0\u00a0哈哈\u00a0哈\u00a0\u00a0\u00a0哈哈"),
		},
		{
			[]byte("\u00a0\u00a0\u00a0哈\u00a0哈哈\u00a0\u00a0\u00a0a"),
			[]byte("a\u00a0\u00a0\u00a0哈哈\u00a0哈\u00a0\u00a0\u00a0"),
		},
		{
			[]byte("z\u00a0\u00a0\u00a0哈"),
			[]byte("哈\u00a0\u00a0\u00a0z"),
		},
		{
			[]byte("z\u00a0\u00a0\u00a0"),
			[]byte("\u00a0\u00a0\u00a0z"),
		},
		{
			[]byte("\u00a0\u00a0\u00a0"),
			[]byte("\u00a0\u00a0\u00a0"),
		},
		{
			[]byte("\u00a0"),
			[]byte("\u00a0"),
		},
		{
			[]byte("hahah\u00a0hahahahahaaa\u00a0\u00a0"),
			[]byte("\u00a0\u00a0aaahahahahah\u00a0hahah"),
		},
	}

	sliceEqual := func(lhs, rhs []byte) bool {
		if len(lhs) != len(rhs) {
			return false
		}

		for i := 0; i < len(lhs); i++ {
			if lhs[i] != rhs[i] {
				return false
			}
		}
		return true
	}

	for _, test := range tests {
		got := make([]byte, len(test.input))
		copy(got, test.input)
		got = reverse(got)

		if !sliceEqual(got, test.want) {
			t.Errorf("replace(%q) = %q, expect %q", test.input, got, test.want)
			t.Errorf("replace(%q) = %v, expect %v", test.input, got, test.want)
		}
	}
}
