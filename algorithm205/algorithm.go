package algorithm205

func isIsomorphic(s string, t string) bool {
	m1, m2 := make(map[uint8]uint8, len(s)), make(map[uint8]uint8, len(t))
	for i := range s {
		c1, ok1 := m1[s[i]]
		c2, ok2 := m2[t[i]]
		if !ok1 && ok2 || ok1 && !ok2 {
			return false
		}
		if !ok1 && !ok2 {
			m1[s[i]], m2[t[i]] = t[i], s[i]
			continue
		}
		if c1 != t[i] || c2 != s[i] {
			return false
		}
	}
	return true
}
