package algorithm125

import "strings"

func isPalindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}
	s = strings.ToLower(s)
	for i, j := 0, n-1; i < j; {
		for ; i < j && !isValidChar(s[i]); i++ {
		}
		for ; i < j && !isValidChar(s[j]); j-- {
		}
		if s[i] != s[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

func isValidChar(c uint8) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}
