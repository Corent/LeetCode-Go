package algorithm187

func findRepeatedDnaSequences(s string) []string {
	ans, set, m := make([]string, 0), make(map[string]struct{}), make(map[int]int)
	vals, key, n := map[uint8]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}, 0, len(s)
	for i := 0; i < n; i++ {
		key = ((key << 2) | vals[s[i]]) & 0x000fffff
		if i < 9 {
			continue
		}
		cnt, ok := m[key]
		if !ok {
			m[key] = 1
			continue
		}
		m[key] = cnt + 1
		str := s[i-9 : i+1]
		if _, ok = set[str]; !ok {
			ans = append(ans, str)
			set[str] = struct{}{}
		}
	}
	return ans
}
