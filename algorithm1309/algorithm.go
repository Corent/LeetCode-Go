package algorithm1309

import "strings"

func freqAlphabets(s string) string {
	builder := strings.Builder{}
	for idx, n := 0, len(s); idx < n; idx++ {
		if idx+2 < n && s[idx+2] == '#' {
			builder.WriteByte('a' - 1 + (s[idx]-'0')*10 + (s[idx+1] - '0'))
			idx += 2
		} else {
			builder.WriteByte('a' - 1 + (s[idx] - '0'))
		}
	}
	return builder.String()
}
