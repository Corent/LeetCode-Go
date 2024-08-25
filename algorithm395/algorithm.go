package algorithm395

func longestSubstring(s string, k int) int {
	rst := 0
	for unique := 1; unique <= 26; unique++ {
		uniqueChInWin, uniqueChKInWin, cnts := 0, 0, [26]int{}
		for left, right := 0, 0; right < len(s); right++ {
			idx := s[right] - 'a'
			if cnts[idx] == 0 {
				uniqueChInWin++
			}
			if cnts[idx]++; cnts[idx] == k {
				uniqueChKInWin++
			}
			if uniqueChInWin > unique {
				for left < right {
					leftIdx := s[left] - 'a'
					if cnts[leftIdx]--; cnts[leftIdx] == 0 {
						uniqueChInWin--
					}
					if cnts[leftIdx] == k-1 {
						uniqueChKInWin--
					}
					left++
					if uniqueChInWin == unique {
						break
					}
				}
			}
			if uniqueChInWin == unique && uniqueChKInWin == unique && right-left+1 > rst {
				rst = right - left + 1
			}
		}
	}
	return rst
}
