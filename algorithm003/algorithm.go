package algorithm003

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	lastPos := make([]int, 256)
	for i := range lastPos {
		lastPos[i] = -1
	}
	maxLen, lastRepeatPos := 0, -1
	for i := range s {
		if lp := lastPos[s[i]]; lp > lastRepeatPos {
			lastRepeatPos = lp
		}
		if i-lastRepeatPos > maxLen {
			maxLen = i - lastRepeatPos
		}
		lastPos[s[i]] = i
	}
	return maxLen
}
