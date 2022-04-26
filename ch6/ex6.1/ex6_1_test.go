package ex6_1

import (
	"gopl-exercise/ch6/intset"
	"testing"
)

func cast(s *IntSet) *intset.IntSet {
	return (*intset.IntSet)(s)
}

func TestIntSetLen(t *testing.T) {
	var s IntSet

	check := func(res int, exp int) {
		if n := s.Len(); n != exp {
			t.Errorf("s.Len() = %d, expect %d\n", n, exp)
		}
	}

	check(s.Len(), 0)
	num := 9999
	for i := 0; i != num; i++ {
		cast(&s).Add(i)
	}
	check(s.Len(), num)
}

func TestIntSetRemove(t *testing.T) {
	var s IntSet

	check := func(res int, exp int) {
		if n := s.Len(); n != exp {
			t.Errorf("s.Len() = %d, expect %d\n", n, exp)
		}
	}

	check(s.Len(), 0)
	num := 9999
	for i := 0; i != num; i++ {
		cast(&s).Add(i)
	}
	check(s.Len(), num)

	for i := 0; i != num; i++ {
		s.Remove(i)
		check(s.Len(), num-i-1)
		if cast(&s).Has(i) {
			t.Errorf("s.Has(%d) = True, expect False\n", i)
		}
	}
}

func TestIntSetClear(t *testing.T) {
	var s IntSet

	check := func(res int, exp int) {
		if n := s.Len(); n != exp {
			t.Errorf("s.Len() = %d, expect %d\n", n, exp)
		}
	}

	check(s.Len(), 0)
	num := 9999
	for i := 0; i != num; i++ {
		cast(&s).Add(i)
	}
	check(s.Len(), num)

	s.Clear()
	check(s.Len(), 0)
}

func TestIntSetCopy(t *testing.T) {
	var s IntSet
	p := s.Copy()

	if p == &s {
		t.Errorf("The memory address of p and s is same\n")
	}

	check := func(set *IntSet, res int, exp int) {
		if n := set.Len(); n != exp {
			t.Errorf("set.Len() = %d, expect %d\n", n, exp)
		}
	}

	check(&s, s.Len(), 0)
	num := 9999
	for i := 0; i != num; i++ {
		cast(&s).Add(i)
	}
	check(&s, s.Len(), num)

	check(p, p.Len(), 0)
	for i := 0; i != num*2; i++ {
		cast(p).Add(i)
	}

	check(&s, s.Len(), num)
	check(p, p.Len(), num*2)
}
