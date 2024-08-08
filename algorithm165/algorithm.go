package algorithm165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	ss1 := strings.Split(strings.TrimSpace(version1), ".")
	ss2 := strings.Split(strings.TrimSpace(version2), ".")
	for i, j := 0, 0; i < len(ss1) || j < len(ss2); i, j = i+1, j+1 {
		var n1, n2 int
		if i < len(ss1) {
			n1, _ = strconv.Atoi(ss1[i])
		}
		if j < len(ss2) {
			n2, _ = strconv.Atoi(ss2[j])
		}
		if n1 > n2 {
			return 1
		} else if n1 < n2 {
			return -1
		}
	}
	return 0
}
