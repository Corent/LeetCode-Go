package algorithm032

import "strings"

func longestValidParentheses(s string) int {
	if len(s) <= 1 || strings.Index(s, ")") == -1 {
		return 0
	}
	ans, start := 0, -1
	stack := make([]int, 0)
	for i, c := range s {
		if c == '(' { // 碰到'('无条件压栈
			stack = append(stack, i)
			continue
		}
		if len(stack) == 0 { // 碰到')'且栈空，更新合法括号对的起点start
			start = i
			continue
		}
		idx := start
		stack = stack[:len(stack)-1] // 碰到')'且栈不空，计算当前合法括号对的长度
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}
		if deep := i - idx; deep > ans {
			ans = deep
		}
	}
	return ans
}
