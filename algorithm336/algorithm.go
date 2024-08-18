package algorithm336

func palindromePairs(words []string) [][]int {
	if len(words) < 2 {
		return [][]int{}
	}
	m, rst := make(map[string]int), make([][]int, 0)
	for i, word := range words {
		m[word] = i
	}
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			left, right := word[:j], word[j:]
			if isPalindrome(left) {
				reversedRight := reverse(right)
				if idx, ok := m[reversedRight]; ok && idx != i {
					rst = append(rst, []int{idx, i})
				}
			}
			if isPalindrome(right) {
				reversedLeft := reverse(left)
				if idx, ok := m[reversedLeft]; ok && idx != i && len(right) > 0 {
					rst = append(rst, []int{i, idx})
				}
			}
		}
	}
	return rst
}

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	if len(s) < 2 {
		return s
	}
	n := len(s)
	chs := make([]uint8, n)
	for i := 0; i <= n/2; i++ {
		chs[i], chs[n-1-i] = s[n-1-i], s[i]
	}
	return string(chs)
}
