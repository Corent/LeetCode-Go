package algorithm479

import (
	"math"
	"strconv"
)

func largestPalindrome(n int) int {
	upper := math.Pow10(n) - 1
	lower := upper / 10
	for i := upper; i > lower; i-- {
		t := strconv.Itoa(int(i))
		p, _ := strconv.Atoi(t + reverse(t))
		for j := int(upper); j*j > p; j-- {
			if p%j == 0 {
				return p % 1337
			}
		}
	}
	return 9
}

func reverse(s string) string {
	if len(s) < 2 {
		return s
	}
	bytes := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
