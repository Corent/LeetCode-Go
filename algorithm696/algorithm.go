package algorithm696

func countBinarySubstrings(s string) int {
	groups := make([]int, 0)
	m, n := s[0], 1
	for i := 1; i <= len(s); i++ {
		if i == len(s) {
			groups = append(groups, n)
			break
		}
		if s[i] == m {
			n++
		} else {
			groups, m, n = append(groups, n), s[i], 1
		}
	}
	cnt := 0
	for i := 1; i < len(groups); i++ {
		cnt += minInt(groups[i], groups[i-1])
	}
	return cnt
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
