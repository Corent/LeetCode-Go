package algorithm434

import "strings"

func countSegments(s string) int {
	ss := strings.Split(strings.TrimSpace(s), " ")
	cnt := 0
	for _, str := range ss {
		if len(str) > 0 {
			cnt++
		}
	}
	return cnt
}
