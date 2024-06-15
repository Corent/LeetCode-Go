package algorithm022

var (
	chs []uint8
	ans []string
)

func generateParenthesis(n int) []string {
	if n <= 0 {
		return ans
	}
	chs, ans = make([]uint8, n*2), make([]string, 0)
	helper(0, n, n)
	return ans
}

func helper(idx, left, right int) {
	if idx == len(chs) {
		ans = append(ans, string(chs))
		return
	}
	if left > 0 {
		chs[idx] = '('
		helper(idx+1, left-1, right)
	}
	if left < right {
		chs[idx] = ')'
		helper(idx+1, left, right-1)
	}
	chs[idx] = 0
}
