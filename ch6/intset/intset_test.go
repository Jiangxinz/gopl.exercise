package intset

import (
	"fmt"
	"testing"
)

func TestIntSet(t *testing.T) {
	set := new(IntSet)
	val := 1289
	if set.Has(val) {
		t.Errorf("IntSet.Has(%d) = True, expect False\n", val)
	}
	set.Add(val)
	if !set.Has(val) {
		t.Errorf("IntSet.Has(%d) = False, expect True\n", val)
	}
	fmt.Println(set.Words())
	fmt.Println(set)
}
