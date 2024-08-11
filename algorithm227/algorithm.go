package algorithm227

func calculate(s string) int {
	op, ans, num := uint8('+'), 0, 0
	if len(s) == 0 {
		return ans
	}
	stack := make([]int, 0)
	for i, n := 0, len(s); i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}
		if s[i] < '0' && s[i] != ' ' || i == n-1 {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				m := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, m*num)
			case '/':
				m := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, m/num)
			}
			op, num = s[i], 0
		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		ans += stack[i]
	}
	return ans
}
