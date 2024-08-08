package algorithm168

import "strings"

var chs = []uint8{'Z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y'}

func convertToTitle(columnNumber int) string {
	stack := make([]uint8, 0)
	for columnNumber != 0 {
		stack = append(stack, chs[columnNumber%26])
		columnNumber = (columnNumber - 1) / 26
	}
	builder := strings.Builder{}
	for len(stack) > 0 {
		builder.WriteByte(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return builder.String()
}
