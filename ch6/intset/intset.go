package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Words() *[]uint64 {
	return &(s.words)
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, val := range t.words {
		if i >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[i] |= val
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("{ ")
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for bit := 0; bit < 64; bit++ {
			if word&(1<<uint64(bit)) != 0 {
				fmt.Fprintf(&buf, "%d ", 64*i+bit)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// func (s *IntSet) Copy() *IntSet {
// 				words := make([]uint64, len(s.words))
// 				copy(words, s.words)
// 				return &IntSet{words}
// }
