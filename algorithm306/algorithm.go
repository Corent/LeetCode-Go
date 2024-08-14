package algorithm306

import (
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	n := len(num)
	for i := 1; i < n-1; i++ {
		a, ok1 := convertStrToInt64(num[0:i])
		if !ok1 {
			continue
		}
		for j := i + 1; j < n; j++ {
			b, ok2 := convertStrToInt64(num[i:j])
			if !ok2 {
				continue
			}
			if check(a, b, num[j:]) {
				return true
			}
		}
	}
	return false
}

func convertStrToInt64(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	if len(s) > 1 && s[0] == '0' {
		return 0, false
	}
	v, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return 0, false
	}
	return v, true
}

func check(a, b int64, s string) bool {
	num1, num2 := a, b
	nextNums, nextNum := s, strconv.FormatInt(a+b, 10)
	for strings.HasPrefix(nextNums, nextNum) {
		nextNums = nextNums[len(nextNum):]
		num1, num2 = num2, num1+num2
		nextNum = strconv.FormatInt(num1+num2, 10)
	}
	return len(nextNums) == 0
}
