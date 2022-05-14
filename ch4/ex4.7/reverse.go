package reverse

import "fmt"

func reverse(arr []byte) []byte {
	// 128 -> b1000 0000
	// 192 -> b1100 0000
	// 224 -> b1110 0000
	// 240 -> b1111 0000
	patterns := [...][]uint8{
		// 4字节: 1111 0xxx 10xx xxxx 10xx xxxx 10xx xxxx
		// 11110000 10000000 10000000 10000000
		[]uint8{240, 128, 128, 128},

		// 3字节: 1110 xxxx 10xx xxxx 10xx xxxx
		// 11100000 10000000 10000000
		[]uint8{224, 128, 128},

		// 2字节: 110x xxxx 10xx xxxx
		// 11000000 10000000
		[]uint8{192, 128},

		// 1字节: 0xxx xxxx
		// 00000000
		[]uint8{0},
	}
	masks := [...][]uint8{
		// 11111000 11000000 11000000 11000000
		[]uint8{248, 192, 192, 192},

		// 11110000 11000000 11000000
		[]uint8{240, 192, 192},

		// 11100000 11000000
		[]uint8{224, 192},

		// 10000000
		[]uint8{128},
	}
	i := 0
	for i < len(arr) {
		find := true
		n := 0
		for j, mask := range masks {
			// start match ${len(mask)} byte
			n = 0
			find = true
			for k, val := range mask {
				if (uint8(arr[i+n]) & val) != patterns[j][k] {
					find = false
					break
				}
				n++
			}
			if find {
				break
			}
		}
		if !find {
			panic("impossible: invalid val")
		}
		helper(arr[i : i+n])
		i += n
	}
	helper(arr)
	return arr
}

func helper(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
