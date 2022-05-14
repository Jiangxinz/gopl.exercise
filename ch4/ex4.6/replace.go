package replace

import "unicode/utf8"

// import "fmt"

func replace(arr []byte) []byte {
	length := len(arr)
	if length == 0 {
		return arr
	}

	indexSize := make(map[int]int)
	runeMap := make(map[int]rune)

	s := string(arr)
	for i := 0; i < length; {
		r, size := utf8.DecodeRuneInString(s[i:])
		runeMap[i] = r
		indexSize[i] = size
		i += size
	}

	left, right := -1, 0
	indexSize[-1] = 1
	// fmt.Println("indexSize: ", indexSize)

	for right < length {
		size := indexSize[right]
		r := runeMap[right]
		cnt := 0
		// 如果以\u00a0结尾
		ok := true
		for r == '\u00a0' && right < length {
			cnt++
			right += size
			size, ok = indexSize[right]
			// assert(ok)
			r, ok = runeMap[right]
			// assert(ok)
		}
		if cnt >= 2 {
			size, ok = indexSize[left]
			assert(ok)
			left += size
			assert(left <= right)
			arr[left] = ' '
			indexSize[left] = 1
		} else if cnt == 1 {
			// left += indexSize[left]
			// fmt.Printf("cnt = 1 before: [%d, %d]\n", left, right)
			// FIXME: magic number
			right -= 2
			// fmt.Printf("cnt = 1 after: [%d, %d]\n", left, right)
		}
		size, ok = indexSize[left]
		assert(ok)
		left += size

		// fmt.Printf("[%c, %c], [%d, %d], cnt = %d\n", runeMap[left], runeMap[right], left, right, cnt)
		assert(left <= right)

		size, ok = indexSize[right]
		// assert(ok)
		copy(arr[left:left+size], arr[right:right+size])
		right += size
		// fmt.Printf("[%d, %d]\n", left, right)
		indexSize[left] = size
	}
	// fmt.Printf("left = %d, index_size = %d\n", left, indexSize[left])
	return arr[:left+indexSize[left]]
}

func assert(exp interface{}) {
	val, _ := exp.(bool)
	if !val {
		panic("assertion failed")
	}
}
