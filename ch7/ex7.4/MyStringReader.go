package mystringreader

import (
	"bufio"
	// "fmt"
	"io"
	"strings"
)

type MyReader struct {
	s    string
	scan *bufio.Scanner
}

// html.Parse通常会先传入固定大小的缓冲区，当前测试len和cap均为4096,
// 我觉得因为是拷贝传参，应该避免重分配, 所以应该从cap大小开始读取
// 并且读取时使用索引而不是append函数, 因为p的长度不为0
func (r *MyReader) Read(p []byte) (n int, err error) {
	// fmt.Println(len(p), cap(p))
	for cnt := cap(p); cnt != 0; cnt-- {
		if !r.scan.Scan() {
			err = io.EOF
			break
		}
		// 这里断言传入的p初始内容均为0，即为无效内容
		p[n] = r.scan.Bytes()[0]
		n++
	}
	return n, err
}

func MyStringReader(s string) (r io.Reader, err error) {
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanBytes)
	return &MyReader{s, scan}, nil
}
