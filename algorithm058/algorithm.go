package algorithm058

func lengthOfLastWord(s string) int {
	ans := 0
	if len(s) == 0 {
		return ans
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		j := i - 1
		for ; j >= 0 && s[j] != ' '; j-- {
		}
		return i - j
	}
	return ans
}
