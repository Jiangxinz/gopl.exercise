package ex6_2

import (
	"gopl-exercise/ch6/intset"
)

type IntSet intset.IntSet

func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		(*intset.IntSet)(s).Add(val)
	}
}
