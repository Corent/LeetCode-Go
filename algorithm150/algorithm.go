package algorithm150

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+":
			m, n := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], m+n)
		case "-":
			m, n := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], m-n)
		case "*":
			m, n := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], m*n)
		case "/":
			m, n := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], m/n)
		default:
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)
		}
	}
	return stack[0]
}
