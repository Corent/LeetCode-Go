package algorithm005

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	m, maxLen, ans := len(s), 0, ""
	isPalindrome := make([][]bool, m)
	for i := m - 1; i >= 0; i-- {
		isPalindrome[i] = make([]bool, m)
		for j := i; j < m; j++ {
			l := j - i + 1
			if s[i] == s[j] && (l <= 3 || isPalindrome[i+1][j-1]) {
				isPalindrome[i][j] = true
				if l > maxLen {
					maxLen, ans = l, s[i:j+1]
				}
			}
		}
	}
	return ans
}
