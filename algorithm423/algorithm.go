package algorithm423

import (
	"strconv"
	"strings"
)

func originalDigits(s string) string {
	letters, numbers := make([]int, 26), make([]int, 10)
	for _, c := range s {
		letters[c-'a']++
	}
	numbers[8] = letters['g'-'a']
	numbers[6] = letters['x'-'a']
	numbers[4] = letters['u'-'a']
	numbers[2] = letters['w'-'a']
	numbers[0] = letters['z'-'a']
	numbers[5] = letters['f'-'a'] - numbers[4]
	numbers[3] = letters['h'-'a'] - numbers[8]
	numbers[7] = letters['s'-'a'] - numbers[6]
	numbers[1] = letters['o'-'a'] - numbers[0] - numbers[2] - numbers[4]
	numbers[9] = letters['i'-'a'] - numbers[5] - numbers[6] - numbers[8]
	builder := strings.Builder{}
	for i := range numbers {
		for numbers[i] > 0 {
			builder.WriteString(strconv.Itoa(i))
			numbers[i]--
		}
	}
	return builder.String()
}
