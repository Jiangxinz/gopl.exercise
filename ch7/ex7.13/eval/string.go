package eval

import (
	"bytes"
	"fmt"
	"strconv"
)

var prefixStk = []string{""}
var isFirstStk = []bool{false}

func stringUnaryNode(t string, v string) string {
	var buf bytes.Buffer
	prefix := prefixStk[len(prefixStk)-1]
	buf.WriteString(prefix)
	buf.WriteString("└──")
	buf.WriteString(fmt.Sprintf("(%s: %s)\n", t, v))
	return buf.String()
}

// Var
func (v Var) String() string {
	return stringUnaryNode("Var", string(v))
}

// literal
func (l literal) String() string {
	return stringUnaryNode("literal", strconv.FormatFloat(float64(l), 'E', -1, 64))
}

// TODO: eliminates redundant code
// Unary
func (u unary) String() string {
	var buf bytes.Buffer
	prefix := prefixStk[len(prefixStk)-1]
	isFirst := isFirstStk[len(isFirstStk)-1]

	buf.WriteString(prefix)
	buf.WriteString("└──")
	buf.WriteString(fmt.Sprintf("(unary: %q)\n", u.op))

	// if isFirst {
	// 	prefix += "│   "
	// } else {
	// 	prefix += "    "
	// }
	prefix += "    "
	isFirst = true
	// push
	prefixStk = append(prefixStk, prefix)
	isFirstStk = append(isFirstStk, isFirst)
	buf.WriteString(u.x.String())
	// pop
	prefixStk = prefixStk[:len(prefixStk)-1]
	isFirstStk = isFirstStk[:len(isFirstStk)-1]
	return buf.String()
}

// Binary
func (b binary) String() string {
	var buf bytes.Buffer
	prefix := prefixStk[len(prefixStk)-1]
	isFirst := isFirstStk[len(isFirstStk)-1]

	buf.WriteString(prefix)
	if isFirst {
		buf.WriteString("├──")
	} else {
		buf.WriteString("└──")
	}
	buf.WriteString(fmt.Sprintf("(binary: %q)\n", b.op))

	if isFirst {
		prefix += "│   "
	} else {
		prefix += "    "
	}
	isFirst = true
	// push
	prefixStk = append(prefixStk, prefix)
	isFirstStk = append(isFirstStk, isFirst)
	buf.WriteString(b.x.String())
	// pop
	prefixStk = prefixStk[:len(prefixStk)-1]
	isFirstStk = isFirstStk[:len(isFirstStk)-1]

	prefix = prefixStk[len(prefixStk)-1]
	isFirst = isFirstStk[len(isFirstStk)-1]
	if isFirst {
		prefix += "│   "
	} else {
		prefix += "    "
	}
	isFirst = false
	// push
	prefixStk = append(prefixStk, prefix)
	isFirstStk = append(isFirstStk, isFirst)
	buf.WriteString(b.y.String())
	// pop
	prefixStk = prefixStk[:len(prefixStk)-1]
	isFirstStk = isFirstStk[:len(isFirstStk)-1]

	return buf.String()
}

// call
func (c call) String() string {
	var buf bytes.Buffer
	prefix := prefixStk[len(prefixStk)-1]
	isFirst := isFirstStk[len(isFirstStk)-1]

	buf.WriteString(prefix)
	if isFirst {
		buf.WriteString("├──")
	} else {
		buf.WriteString("└──")
	}
	buf.WriteString(fmt.Sprintf("(call: %q)\n", c.fn))

	if len(c.args) == 0 {
		return buf.String()
	}

	if isFirst {
		prefix += "│   "
	} else {
		prefix += "    "
	}
	isFirst = true
	// push
	prefixStk = append(prefixStk, prefix)
	isFirstStk = append(isFirstStk, isFirst)
	buf.WriteString(c.args[0].String())
	// pop
	prefixStk = prefixStk[:len(prefixStk)-1]
	isFirstStk = isFirstStk[:len(isFirstStk)-1]

	for i := 1; i < len(c.args); i++ {
		prefix = prefixStk[len(prefixStk)-1]
		isFirst = isFirstStk[len(isFirstStk)-1]
		if isFirst {
			prefix += "│   "
		} else {
			prefix += "    "
		}
		isFirst = false
		// push
		prefixStk = append(prefixStk, prefix)
		isFirstStk = append(isFirstStk, isFirst)
		buf.WriteString(c.args[i].String())
		// pop
		prefixStk = prefixStk[:len(prefixStk)-1]
		isFirstStk = isFirstStk[:len(isFirstStk)-1]
	}

	return buf.String()
}
