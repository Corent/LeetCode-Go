package algorithm151

import (
	"regexp"
	"strings"
)

var chs []uint8

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	reg, _ := regexp.Compile("\\s+")
	s = reg.ReplaceAllString(s, " ")
	n := len(s)
	if n < 2 {
		return s
	}
	chs = []byte(s)
	for i := 0; i < n; {
		for i < n && chs[i] == ' ' {
			i++
		}
		if i == n {
			break
		}
		var j = i + 1
		for j < n && chs[j] != ' ' {
			j++
		}
		reverse(i, j-1)
		i = j
	}
	reverse(0, n-1)
	return string(chs)
}

func reverse(from, to int) {
	for from < to {
		chs[from], chs[to] = chs[to], chs[from]
		from, to = from+1, to-1
	}
}
