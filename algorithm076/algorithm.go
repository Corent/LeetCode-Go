package algorithm076

import "math"

func minWindow(s string, t string) string {
	ans, tHash := "", make([]int, 256)
	for _, c := range t {
		tHash[c]++
	}
	ansLen, tLen, tInWinCnt := math.MaxInt, len(t), 0
	for winStart, winEnd := 0, 0; winEnd < len(s); winEnd++ {
		if tHash[s[winEnd]] > 0 {
			tInWinCnt++
		}
		tHash[s[winEnd]]--
		for tInWinCnt >= tLen {
			winLen := winEnd - winStart + 1
			if ansLen > winLen {
				ans, ansLen = s[winStart:winEnd+1], winEnd-winStart+1
			}
			tHash[s[winStart]]++
			if tHash[s[winStart]] > 0 {
				tInWinCnt--
			}
			winStart++
		}
	}
	return ans
}
