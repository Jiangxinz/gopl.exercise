package unique

func unique(arr []string) []string {
	length := len(arr)
	if length == 0 {
		return arr
	}
	i := 0
	for j := 1; j < length; j++ {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
	}
	return arr[:i+1]
}
