package algorithm065

import "strings"

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	s = strings.TrimSpace(s)
	i, n := 0, len(s)
	if s[i] == '+' || s[i] == '-' {
		i++
	}
	ans, dot, exp := false, false, false
	for i < n {
		c := s[i]
		if c >= '0' && c <= '9' {
			ans = true
		} else if c == '.' {
			if dot || exp {
				return false
			}
			dot = true
		} else if c == 'e' || c == 'E' {
			if !ans || exp {
				return false
			}
			ans, exp = false, true
		} else if c == '+' || c == '-' {
			if s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false
		}
		i++
	}
	return ans
}
