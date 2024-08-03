package algorithm078

import "slices"

var (
	current []int
	ans     [][]int
)

func subsets(nums []int) [][]int {
	current, ans = make([]int, 0), make([][]int, 0)
	pick(nums, 0)
	return ans
}

func pick(nums []int, from int) {
	if from == len(nums) {
		ans = append(ans, slices.Clone[[]int](current))
		return
	}
	pick(nums, from+1)
	current = append(current, nums[from])
	pick(nums, from+1)
	current = current[:len(current)-1]
}
