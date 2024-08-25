package algorithm388

import "strings"

func lengthLongestPath(input string) int {
	stack, ss, rst, cnt := make([]string, 0), strings.Split(input, "\n"), 0, 0
	for _, s := range ss {
		for len(stack) > 0 && level(s) <= level(stack[len(stack)-1]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cnt -= len(top) - level(top) + 1
		}
		stack = append(stack, s)
		cnt += len(s) - level(s) + 1
		if n := cnt - 1; strings.Contains(s, ".") && n > rst {
			rst = n
		}
	}
	return rst
}

func level(s string) int {
	sum := 0
	for i := 0; i < len(s) && s[i] == '\t'; i++ {
		sum++
	}
	return sum
}
