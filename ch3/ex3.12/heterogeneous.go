package heterogeneous

func check(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := make(map[rune]int)

	for _, r := range s1 {
		count[r]++
	}

	for _, r := range s2 {
		ref, ok := count[r]
		if !ok || ref == 0 {
			return false
		}
		count[r]--
	}
	return true
}
