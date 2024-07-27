package algorithm051

var (
	N    int
	ans  [][]string
	nums []int   // 每层Q的位置
	ch   []uint8 // 每一层的字符串
)

func solveNQueens(n int) [][]string {
	N, ans, nums, ch = n, make([][]string, 0), make([]int, n), make([]uint8, n)
	for i := 0; i < n; i++ {
		nums[i], ch[i] = i, '.'
	}
	solveNQueensHelper(0)
	return ans
}

func solveNQueensHelper(level int) {
	if level == N {
		tmp := make([]string, N)
		for i := 0; i < N; i++ {
			ch[nums[i]] = 'Q'
			tmp[i] = string(ch)
			ch[nums[i]] = '.'
		}
		ans = append(ans, tmp)
		return
	}
	for i := 0; i < N; i++ {
		if !check(i, level) {
			continue
		}
		nums[level] = i
		solveNQueensHelper(level + 1)
		nums[level] = 0
	}
}

func check(n, level int) bool { // 检查第level层的第n个位置是否可以放置
	for i := 0; i < level; i++ {
		if nums[i] == n || i-level == nums[i]-n || level-i == nums[i]-n {
			return false
		}
	}
	return true
}
