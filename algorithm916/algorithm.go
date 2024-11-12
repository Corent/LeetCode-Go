package algorithm916

func wordSubsets(words1 []string, words2 []string) []string {
	rst := make([]string, 0)
	BHash := [26]int{}
	for _, word := range words2 {
		BHashTmp := [26]int{}
		for _, w := range word {
			BHashTmp[w-'a']++
		}
		for i := range BHash {
			BHash[i] = maxInt(BHash[i], BHashTmp[i])
		}
	}
	for _, word := range words1 {
		AHash := [26]int{}
		for _, w := range word {
			AHash[w-'a']++
		}
		idx := 0
		for idx < 26 {
			if AHash[idx] < BHash[idx] {
				break
			}
			idx++
		}
		if idx == 26 {
			rst = append(rst, word)
		}
	}
	return rst
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}
