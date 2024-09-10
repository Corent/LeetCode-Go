package algorithm405

import "strings"

var hex = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f",
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	stack, cnt := make([]string, 0), 0
	for num != 0 && cnt < 8 {
		stack = append(stack, hex[num&0xf])
		num, cnt = num>>4, cnt+1
	}
	builder := &strings.Builder{}
	for len(stack) > 0 {
		builder.WriteString(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return builder.String()
}
