package replace

import "testing"

// import "fmt"

func TestReplace(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{
			[]byte("哈哈\u00a0\u00a0哈\u00a0哈哈\u00a0\u00a0a"),
			[]byte("哈哈 哈\u00a0哈哈 a"),
		},
		{
			[]byte("哈哈\u00a0\u00a0\u00a0哈\u00a0哈哈\u00a0\u00a0\u00a0a"),
			[]byte("哈哈 哈\u00a0哈哈 a"),
		},
		{
			[]byte("\u00a0\u00a0\u00a0哈\u00a0哈哈\u00a0\u00a0\u00a0a"),
			[]byte(" 哈\u00a0哈哈 a"),
		},
		{
			[]byte("z\u00a0\u00a0\u00a0哈"),
			[]byte("z 哈"),
		},
		{
			[]byte("z\u00a0\u00a0\u00a0"),
			[]byte("z "),
		},
		{
			[]byte("\u00a0\u00a0\u00a0"),
			[]byte(" "),
		},
		{
			[]byte("\u00a0"),
			[]byte("\u00a0"),
		},
		{
			[]byte("hahah\u00a0hahahahahaaa\u00a0\u00a0"),
			[]byte("hahah\u00a0hahahahahaaa "),
		},
	}

	// s := fmt.Sprintf("空 格,空\u00a0格, 空 格")
	// for _, r := range s {
	// 	fmt.Printf("%c, %[1]d\n", rune(r), r)
	// }

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
		got = replace(got)

		if !sliceEqual(got, test.want) {
			t.Errorf("replace(%q) = %q, expect %q", test.input, got, test.want)
			t.Errorf("replace(%q) = %v, expect %v", test.input, got, test.want)
		}
	}
}
