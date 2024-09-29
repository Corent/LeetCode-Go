package algorithm459

func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}
	n := len(s)
	for i := 1; i < n; i++ {
		if n%i != 0 {
			continue
		}
		s0, j := s[:i], i
		for ; j < n; j += i {
			if s[j:j+i] != s0 {
				break
			}
		}
		if j == n {
			return true
		}
	}
	return false
}
