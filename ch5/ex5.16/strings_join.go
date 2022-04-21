package main

import (
	"fmt"
	"os"
)

func Join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var totalLen int
	for _, s := range strs {
		totalLen += len(s)
	}

	totalLen += len(sep) * len(strs)

	ret := make([]byte, 0, totalLen)

	for _, s := range strs {
		ret = appendString(ret, s)
		ret = appendString(ret, sep)
	}

	ret = ret[:len(ret)-len(sep)]

	expectLen := totalLen - len(sep)
	if len(ret) != expectLen {
		fmt.Fprintf(os.Stderr, "len(%q) = %d, expect %d\n", ret, len(ret), expectLen)
		os.Exit(1)
	}

	return string(ret)
}

func appendString(ret []byte, s string) []byte {
	for i := 0; i < len(s); i++ {
		ret = append(ret, s[i])
	}
	return ret
}
