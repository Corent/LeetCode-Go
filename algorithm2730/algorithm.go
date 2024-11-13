package algorithm2730

func findLongestSubStr(s string) int {
	n, rst := len(s), 0
	repeat, start, end := 0, 0, 0
	for end < n {
		if end > 0 && s[end] == s[end-1] {
			repeat++
		}
		for repeat > 1 {
			if s[start] == s[start+1] {
				repeat--
			}
			start++
		}
		rst = maxInt(rst, end-start+1)
		end++
	}
	return rst
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
