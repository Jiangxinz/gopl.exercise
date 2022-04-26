package ex6_4

import ()

func (s *IntSet) Elems() []uint64 {
	var ret []uint64
	for i, word := range s.words {
		for bit := 0; bit != 64; bit++ {
			if word&(1<<uint(bit)) != 0 {
				ret = append(ret, i*64+bit)
			}
		}
	}
	return ret
}
