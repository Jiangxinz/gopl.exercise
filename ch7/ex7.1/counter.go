package counter

import (
	"bufio"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (n int, err error) {
	cnt, _ := count(string(p), bufio.ScanWords)
	*c = LineCounter(cnt)
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (n int, err error) {
	cnt, _ := count(string(p), bufio.ScanLines)
	*c = LineCounter(cnt)
	return len(p), nil
}

func count(s string, f bufio.SplitFunc) (n int, err error) {
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(f)
	ret := 0
	for scan.Scan() {
		ret += 1
	}
	return ret, nil
}
