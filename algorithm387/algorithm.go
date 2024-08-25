package algorithm387

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}
	chs := [26]int{}
	for _, ch := range s {
		chs[ch-'a']++
	}
	for i, ch := range s {
		if chs[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
