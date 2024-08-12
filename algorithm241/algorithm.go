package algorithm241

import "strconv"

func diffWaysToCompute(expression string) []int {
	ans := make([]int, 0)
	if len(expression) == 0 {
		return ans
	}
	for i, c := range expression {
		if c != '+' && c != '-' && c != '*' {
			continue
		}
		left := diffWaysToCompute(expression[:i])
		right := diffWaysToCompute(expression[i+1:])
		for _, l := range left {
			for _, r := range right {
				switch c {
				case '+':
					ans = append(ans, l+r)
				case '-':
					ans = append(ans, l-r)
				case '*':
					ans = append(ans, l*r)
				}
			}
		}
	}
	if len(ans) == 0 {
		n, _ := strconv.Atoi(expression)
		ans = append(ans, n)
	}
	return ans
}
