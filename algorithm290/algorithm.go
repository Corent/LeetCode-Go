package algorithm290

import "strings"

func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")
	if len(pattern) != len(ss) {
		return false
	}
	m1, m2 := make(map[rune]string, len(pattern)), make(map[string]rune, len(ss))
	for i, c := range pattern {
		v1, ok1 := m1[c]
		v2, ok2 := m2[ss[i]]
		if ok1 && !ok2 || !ok1 && ok2 {
			return false
		}
		if ok1 && ok2 {
			if v1 != ss[i] || v2 != c {
				return false
			}
			continue
		}
		m1[c], m2[ss[i]] = ss[i], c
	}
	return true
}
