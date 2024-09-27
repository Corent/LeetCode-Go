package algorithm409

func longestPalindrome(s string) int {
	chs, ans := make([]int, 256), 0
	for _, ch := range s {
		if chs[ch]++; chs[ch] == 2 {
			ans, chs[ch] = ans+2, 0
		}
	}
	if ans < len(s) {
		ans++
	}
	return ans
}
