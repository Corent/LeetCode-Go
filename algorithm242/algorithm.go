package algorithm242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chs := [256]int{}
	for _, c := range s {
		chs[c]++
	}
	for _, c := range t {
		if chs[c] == 0 {
			return false
		}
		chs[c]--
	}
	return true
}
