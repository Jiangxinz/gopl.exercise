package limitreader

import (
	"io"
)

type MyLimitReader struct {
	reader io.Reader
	left   int64
}

func (r *MyLimitReader) Read(p []byte) (n int, err error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	for r.left > 0 {
		// var buf = make([]byte, r.left)
		p = p[:r.left]
		nRead, err := r.reader.Read(p)
		if err != nil {
			break
		}
		// copy(p, buf)
		n += nRead
		r.left -= int64(nRead)
	}
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &MyLimitReader{r, n}
}
