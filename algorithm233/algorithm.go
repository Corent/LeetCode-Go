package algorithm233

import (
	"math"
	"strconv"
)

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	cnt, m := 0, len(strconv.Itoa(n))
	for i := 1; i <= m; i++ {
		pow10 := int(math.Pow10(i - 1))
		high := (n / pow10) / 10
		cnt += high * pow10

		digit := (n / pow10) % 10
		if digit > 1 {
			cnt += pow10
		} else if digit == 1 {
			cnt += n%pow10 + 1
		}
	}
	return cnt
}
