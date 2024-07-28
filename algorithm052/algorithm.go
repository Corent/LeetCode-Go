package algorithm052

var (
	N, cnt int
	nums   []int // 每层Q的位置
)

func totalNQueens(n int) int {
	N, cnt, nums = n, 0, make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	totalNQueensHelper(0)
	return cnt
}

func totalNQueensHelper(level int) {
	if level == N {
		cnt++
		return
	}
	for i := 0; i < N; i++ {
		if !check(i, level) {
			continue
		}
		nums[level] = i
		totalNQueensHelper(level + 1)
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
