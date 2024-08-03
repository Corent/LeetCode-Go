package algorithm077

import "slices"

var (
	nums    []int
	current []int
	ans     [][]int
)

func combine(n int, k int) [][]int {
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	current, ans = make([]int, 0), make([][]int, 0)
	pick(0, k)
	return ans
}

func pick(from, k int) {
	if len(nums)-from < k {
		return
	}
	if k == 0 {
		ans = append(ans, slices.Clone[[]int](current))
		return
	}
	for i := from; i < len(nums); i++ {
		current = append(current, nums[i])
		pick(i+1, k-1)
		current = current[:len(current)-1]
	}
}
