package rotate

import "testing"

func TestRotate(t *testing.T) {
	var tests = []struct {
		inputArr []int
		inputN   int
		want     []int
	}{
		{
			[]int{1, 3, 5, 7, 9},
			1,
			[]int{3, 5, 7, 9, 1},
		},
		{
			[]int{1, 3, 5, 7, 9},
			5,
			[]int{1, 3, 5, 7, 9},
		},
		{
			[]int{1, 3, 5, 7, 9},
			6,
			[]int{3, 5, 7, 9, 1},
		},
	}

	sliceEqual := func(lhs, rhs []int) bool {
		len1, len2 := len(lhs), len(rhs)
		if len1 != len2 {
			return false
		}

		for i := 0; i < len1; i++ {
			if lhs[i] != rhs[i] {
				return false
			}
		}
		return true
	}

	for _, test := range tests {
		got := make([]int, len(test.inputArr))
		copy(got, test.inputArr)
		rotate(got, test.inputN)
		if !sliceEqual(got, test.want) {
			t.Errorf("rotate(%v, %d) = %v, expect %v\n", test.inputArr, test.inputN, got, test.want)
		}
	}
}
