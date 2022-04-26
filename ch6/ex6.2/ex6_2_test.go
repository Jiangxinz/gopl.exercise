package ex6_2

import (
	ex6_1 "gopl-exercise/ch6/ex6.1"
	"testing"
)

func TestAddAll(t *testing.T) {
	var s IntSet
	p := (*ex6_1.IntSet)(&s)

	arr1 := []int{1, 3, 5, 7, 9}
	s.AddAll(arr1...)

	if p.Len() != len(arr1) {
		t.Errorf("s.Len() = %d, expect %d\n", p.Len(), len(arr1))
	}
}
