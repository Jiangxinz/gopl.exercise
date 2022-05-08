package comma

import (
	"bytes"
	"fmt"
	"strings"
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
				if !unicode.IsDigit(r) && r != '-' && r != '+' && r != '.' {
					return false
				}
			}
			return true
		},
		// // TODO: use reg-exp check s is valid number
	}

	for _, check := range checks {
		if !check(s) {
			return s, false
		}
	}

	// bias := 0
	sentinel := '?'
	flag := sentinel
	if s[0] == '+' || s[0] == '-' {
		flag = rune(s[0])
		s = s[1:]
	}
	dot := strings.Index(s, ".")
	if dot < 0 {
		if flag != sentinel {
			return string(flag) + helper(s, 3, -1), true
		}
		return helper(s, 3, -1), true
	}
	var buf bytes.Buffer
	// fmt.Printf("input %s, bias = %d, dot = %d, [%s]+[%s]\n", s, bias, dot, s[bias:dot], s[dot:])
	if flag != sentinel {
		buf.WriteString(string(flag))
	}
	buf.WriteString(helper(s[:dot], 3, -1))
	buf.WriteByte('.')
	buf.WriteString(helper(s[dot+1:], 3, 0))
	return buf.String(), true
}

func helper(s string, n int, head int) string {
	if len(s) == 0 {
		return s
	}
	if n < 1 {
		panic(fmt.Sprintf("Invalid parameter n: %d\n", n))
	}
	var buf bytes.Buffer
	if head < 0 {
		head = len(s) % n
	}
	buf.WriteString(s[:head])
	// fmt.Printf("helper: s = %s, len(s) = %d, head = %d\n", s, len(s), head)
	for i := head; i < len(s); i = i + 3 {
		buf.WriteByte(',')
		end := i + 3
		if end > len(s) {
			end = len(s)
		}
		buf.WriteString(s[i:end])
	}
	ret := buf.String()
	if ret[0] == ',' {
		ret = ret[1:]
	}
	// fmt.Printf("helper: ret = %s\n", ret)
	return ret
}
