package main

import (
	"crypto/sha256"
	"fmt"
)

func popCount(val uint8) (ret int) {
	for val != 0 {
		ret++
		val &= (val - 1)
	}
	return
}

func diff(a, b []byte) (ret int) {
	c1 := sha256.Sum256(a)
	c2 := sha256.Sum256(b)
	for i := 0; i < 32; i++ {
		val1, val2 := c1[i], c2[i]
		ret += popCount(val1 ^ val2)
	}
	return
}

func main() {
	a := []byte("x")
	b := []byte("X")

	fmt.Printf("diff %q and %q = %d\n", a, b, diff(a, b))
}
