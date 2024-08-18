package algorithm318

func maxProduct(words []string) int {
	ans := 0
	if len(words) == 2 {
		return ans
	}
	letters := make([]int, len(words))
	for idx, word := range words {
		for _, ch := range word {
			letters[idx] |= 1 << (ch - 'a')
		}
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if letters[i]&letters[j] != 0 {
				continue
			}
			if next := len(words[i]) * len(words[j]); next > ans {
				ans = next
			}
		}
	}
	return ans
}
