package main

func max(vals ...int) int {
	return helper(greater, vals...)
}

func min(vals ...int) int {
	return helper(func(lhs int, rhs int) bool {
		return greater(rhs, lhs)
	}, vals...)
}

func max2(first int, vals ...int) int {
	return max(first, helper(greater, vals...))
}

func min2(first int, vals ...int) int {
	return min(first, helper(func(lhs int, rhs int) bool {
		return greater(rhs, lhs)
	}, vals...))
}

func greater(lhs int, rhs int) bool {
	if lhs > rhs {
		return true
	}
	return false
}

func helper(cmp func(int, int) bool, vals ...int) int {
	if len(vals) == 0 {
		// fmt.Fprintf(os.Stderr, "The number of input is Zero\n")
		// os.Exit(1)
		return 0
	}

	ret := vals[0]
	for i := 1; i < len(vals); i++ {
		if cmp(vals[i], ret) {
			ret = vals[i]
		}
	}
	return ret
}
