package main

import (
	"crypto/sha256"
	// "crypto/sha384"
	"crypto/sha512"
	"flag"
	"fmt"
)

func popCount(val uint8) (ret int) {
	for val != 0 {
		ret++
		val &= (val - 1)
	}
	return
}

func diff(a, b []byte, shaFunc func([]byte) []byte) (ret int) {
	c1 := shaFunc(a)
	c2 := shaFunc(b)
	for i := 0; i < len(c1); i++ {
		val1, val2 := c1[i], c2[i]
		ret += popCount(val1 ^ val2)
	}
	return
}

var sha = flag.String("sha", "sha256", "sha func")

func main() {
	flag.Parse()
	a := []byte("x")
	b := []byte("X")

	mapTable := map[string]func([]byte) []byte{
		"sha256": func(a []byte) []byte {
			ret := sha256.Sum256(a)
			return ret[:]
		},
		// "sha384": func(a []byte) []byte {
		// 	return sha384.Sum384(a)[:]
		// },
		"sha512": func(a []byte) []byte {
			ret := sha512.Sum512(a)
			return ret[:]
		},
	}

	shaFunc, ok := mapTable[*sha]
	if !ok {
		panic(fmt.Sprintf("invalid sha func %s\n", *sha))
	}

	fmt.Printf("diff %q and %q = %d\n", a, b, diff(a, b, shaFunc))
}
