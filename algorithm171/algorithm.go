package algorithm171

func titleToNumber(columnTitle string) int {
	ans, tmp := 0, 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		ans += int(columnTitle[i]-'A'+1) * tmp
		tmp *= 26
	}
	return ans
}
