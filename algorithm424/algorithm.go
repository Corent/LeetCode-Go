package algorithm424

func characterReplacement(s string, k int) int {
	hash, rst, maxCnt := [26]int{}, 0, 0
	for i, j := 0, 0; j < len(s); j++ {
		hash[s[j]-'A']++
		maxCnt = maxInt(maxCnt, hash[s[j]-'A'])
		for j-i+1-maxCnt > k {
			hash[s[i]-'A']--
			i++
		}
		rst = maxInt(rst, j-i+1)
	}
	return rst
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
