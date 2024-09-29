package algorithm443

import (
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	next := 0
	for idx, ch, i := 0, chars[0], 0; i <= len(chars); i++ {
		if i == len(chars) || chars[i] != ch {
			cnt := i - idx
			chars[next], next = ch, next+1
			if cnt > 1 {
				num := strconv.Itoa(cnt)
				for j := 0; j < len(num); j++ {
					chars[next], next = num[j], next+1
				}
			}
			if i < len(chars) {
				idx, ch = i, chars[i]
			}
		}
	}
	return next
}
