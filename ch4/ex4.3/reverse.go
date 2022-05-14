package reverse

const ARRLEN = 5

func reverse(arr *[ARRLEN]int) {
	for i, j := 0, ARRLEN-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
