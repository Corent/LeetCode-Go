package algorithm394

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, string(s[i]))
			continue
		}
		sub := ""
		for len(stack) > 0 && !(len(stack[len(stack)-1]) == 1 && stack[len(stack)-1][0] == '[') {
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sub = c + sub
		}
		stack = stack[:len(stack)-1]
		num := ""
		for len(stack) > 0 && len(stack[len(stack)-1]) == 1 && stack[len(stack)-1][0] >= '0' && stack[len(stack)-1][0] <= '9' {
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num = c + num
		}
		n, _ := strconv.Atoi(num)
		for j := 0; j < n; j++ {
			stack = append(stack, sub)
		}
	}
	return strings.Join(stack, "")
}
