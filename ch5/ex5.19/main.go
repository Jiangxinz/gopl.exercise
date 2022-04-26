package main

import "fmt"

func f() (ret int) {
	type bailout struct{}
	defer func() {
		if p := recover(); (p == bailout{}) {
			ret = 1
		}
	}()

	panic(bailout{})
}

func main() {
	fmt.Println("f() = ", f())
}
