package algorithm214

import "strings"

func shortestPalindrome(s string) string {
	end := len(s) - 1
	for i, j := 0, len(s)-1; i < j; {
		if s[i] == s[j] {
			i, j = i+1, j-1
		} else {
			i, j, end = 0, end-1, end-1
		}
	}
	return reverse(s[end+1:]) + s
}

func reverse(s string) string {
	builder := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteByte(s[i])
	}
	return builder.String()
}
