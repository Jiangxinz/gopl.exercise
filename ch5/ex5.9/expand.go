package main

func expand(s string, f func(string) string) string {
	return helper(s, f, "$foo")
}

func helper(s string, f func(string) string, keyWord string) string {
	keyLen := len(keyWord)
	ret := make([]byte, 0, keyLen)
	j := 0
	for i, _ := range s {
		ret = append(ret, s[i])
		if s[i] == keyWord[j] {
			j++
			if j == keyLen {
				ret = ret[:len(ret)-keyLen]
				str := f(keyWord[1:])
				for k, _ := range str {
					ret = append(ret, str[k])
				}
				j = 0
			}
		} else {
			j = 0
		}
	}
	return string(ret)
}
