package algorithm301

var ans []string

func removeInvalidParentheses(s string) []string {
	l, r := 0, 0
	ans = make([]string, 0)
	for _, c := range s {
		l += chooseIf[int](c == '(', 1, 0)
		if l == 0 {
			r += chooseIf[int](c == ')', 1, 0)
		} else {
			l -= chooseIf[int](c == ')', 1, 0)
		}
	}
	dfs(s, 0, l, r)
	return ans
}

func dfs(s string, start, l, r int) {
	if l == 0 && r == 0 {
		if isValid(s) {
			ans = append(ans, s)
			return
		}
	}

	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] || s[i] != '(' && s[i] != ')' {
			continue
		}
		next := s[:i] + s[i+1:]
		if r > 0 && s[i] == ')' {
			dfs(next, i, l, r-1)
		} else if l > 0 && s[i] == '(' {
			dfs(next, i, l-1, r)
		}
	}
}

func isValid(s string) bool {
	cnt := 0
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if c == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}

func chooseIf[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
