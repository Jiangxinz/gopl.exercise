package comma

import (
	"bytes"
	"fmt"
	"unicode"
)

func Comma(s string) (string, bool) {
	checks := []func(s string) bool{
		func(s string) bool {
			if len(s) == 0 {
				return false
			}
			return true
		},
		func(s string) bool {
			for _, r := range s {
				if !unicode.IsDigit(r) {
					return false
				}
			}
			return true
		},
	}

	for _, check := range checks {
		if !check(s) {
			return s, false
		}
	}
	return helper(s, 3), true
}

func helper(s string, n int) string {
	if n < 1 {
		panic(fmt.Sprintf("Invalid parameter n: %d\n", n))
	}
	var buf bytes.Buffer
	head := len(s) % n
	buf.WriteString(s[:head])
	for i := head; i < len(s); i = i + 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	ret := buf.String()
	if ret[0] == ',' {
		ret = ret[1:]
	}
	return ret
}
