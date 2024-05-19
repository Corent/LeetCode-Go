package algorithm008

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	if len(strings.TrimSpace(s)) == 0 {
		return 0
	}
	sign, start := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == 0x9 {
			start++
		} else {
			break
		}
	}
	if s[start] == '-' {
		sign, start = -1, start+1
	} else if s[start] == '+' {
		start++
	}
	ans := int32(0)
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		n := int32(s[i] - '0')
		if math.MaxInt32/10 >= ans {
			ans *= 10
		} else {
			if ans = math.MaxInt32; sign == -1 {
				ans = math.MinInt32
			}
			return int(ans)
		}

		if math.MaxInt32-n >= ans {
			ans += n
		} else {
			if ans = math.MaxInt32; sign == -1 {
				ans = math.MinInt32
			}
			return int(ans)
		}
	}
	return sign * int(ans)
}
