package algorithm1071

import "regexp"

func gcdOfStrings(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	if m == 0 || n == 0 {
		return ""
	}
	for i, j := minInt(m, n), 1; i > 0; i, j = minInt(m, n)/j, j+1 {
		if m%i == 0 && n%i == 0 {
			s := str1[:i]
			reg, _ := regexp.Compile("^(" + s + ")+$")
			if reg.MatchString(str1) && reg.MatchString(str2) {
				return s
			}
		}
	}
	return ""
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
