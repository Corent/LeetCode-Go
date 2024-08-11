package algorithm224

func calculate(s string) int {
	ans, sign := 0, 1
	if len(s) == 0 {
		return ans
	}
	stack := make([]int, 0)
	for i, n := 0, len(s); i < n; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for i < n && s[i] >= '0' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			ans, i = ans+sign*num, i-1
			continue
		}
		switch s[i] {
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			stack = append(stack, ans)
			stack = append(stack, sign)
			ans, sign = 0, 1
		case ')':
			ans *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}
