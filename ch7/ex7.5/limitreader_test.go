package limitreader

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	type Input struct {
		s string
		n int64
	}
	var tests = []struct {
		Input
		want string
	}{
		{
			Input{
				"123456789",
				4,
			},
			"1234",
		},
		{
			Input{
				"cesar",
				0,
			},
			"",
		},
		{
			Input{
				"",
				0,
			},
			"",
		},
		{
			Input{
				"qwertyuiopasdfghjklzxcvbnm",
				10,
			},
			"qwertyuiop",
		},
	}

	BUFSIZE := 4096

	// assert limitsize <= BUFSIZE

	for _, test := range tests {
		reader := LimitReader(strings.NewReader(test.s), test.n)
		buf := make([]byte, BUFSIZE, BUFSIZE)
		bbuf := buf
		nRead := 0
		for len(bbuf) > 0 {
			n, err := reader.Read(bbuf)
			if err == io.EOF {
				break
			}
			bbuf = bbuf[n:]
			nRead += n
		}
		ans := string(buf[:nRead])
		if test.want != ans {
			t.Errorf(`LimitReader("%s", %d) = "%s", expect "%s"`, test.s, test.n, ans, test.want)
		}
	}
}
