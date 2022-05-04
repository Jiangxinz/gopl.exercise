package ispalindrome

import "testing"

type CheckPalindrome string

func (c CheckPalindrome) Len() int           { return len(c) }
func (c CheckPalindrome) Less(i, j int) bool { return c[i] < c[j] }
func (c CheckPalindrome) Swap(i, j int)      { return }

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"aabaa", true},
		{"aab", false},
		{"", true},
	}

	bool2str := func(val bool) string {
		if val {
			return "true"
		}
		return "false"
	}

	for _, test := range tests {
		if got := isPalindrome(CheckPalindrome(test.input)); got != test.want {
			t.Fatalf("isPalindrome(%s) = %s, expect %s", test.input, bool2str(got), bool2str(test.want))
		}
	}
}
