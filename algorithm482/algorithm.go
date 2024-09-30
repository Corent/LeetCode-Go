package algorithm482

import "strings"

const DIFF = uint8('a' - 'A')

func licenseKeyFormatting(s string, k int) string {
	builder := strings.Builder{}
	for idx, n := len(s)-1, 0; idx >= 0; {
		ch := s[idx]
		idx--
		if ch == '-' {
			continue
		}
		if ch >= 'a' && ch <= 'z' {
			ch -= DIFF
		}
		if n < k {
			n++
		} else {
			n = 1
			builder.WriteByte('-')
		}
		builder.WriteByte(ch)
	}
	return reverse(builder.String())
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
