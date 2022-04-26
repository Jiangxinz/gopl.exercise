package ex6_1

import (
	"fmt"
	"gopl-exercise/ch6/intset"
)

type IntSet intset.IntSet

func (s *IntSet) Len() int {
	cnt := 0
	words := (*intset.IntSet)(s).Words()
	for _, word := range *words {
		for word != 0 {
			cnt++
			word &= (word - 1)
		}
	}
	return cnt
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint64(x%64)
	words := (*intset.IntSet)(s).Words()
	if word >= len(*words) || ((*words)[word]&(1<<bit)) == 0 {
		panic(fmt.Sprintf("%d is not exist in IntSet\n", x))
	}
	(*words)[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	words := (*intset.IntSet)(s).Words()
	*words = make([]uint64, 0)
}

func (s *IntSet) Copy() *IntSet {
	words := (*intset.IntSet)(s).Words()
	// retWords := make([]uint64, len(*words))
	var ret IntSet
	retWords := (*intset.IntSet)(&ret).Words()
	copy(*retWords, *words)
	return &ret
}
