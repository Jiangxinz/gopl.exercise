package countingwriter

import (
	"io"
)

type WrapperWriter struct {
	w   io.Writer
	cnt int64
}

func (ww *WrapperWriter) Write(p []byte) (int, error) {
	n, err := ww.w.Write(p)
	ww.cnt += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var ret = WrapperWriter{
		w,
		0,
	}
	return &ret, &ret.cnt
}
