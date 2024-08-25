package algorithm392

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for i, j := 0, 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			if i++; i == len(s) {
				return true
			}
		}
	}
	return false
}
