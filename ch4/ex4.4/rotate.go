package rotate

func rotate(arr []int, n int) bool {
	length := len(arr)
	if n < 0 || length == 0 {
		return false
	}
	n = n % length
	ret := append(arr, arr[:n]...)
	copy(arr, ret[n:])
	return true
}
