package algorithm216

import "slices"

var (
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans  [][]int
	an   []int
)

func combinationSum3(k int, n int) [][]int {
	ans, an = make([][]int, 0), make([]int, 0)
	combinationSum3Core(k, n, 0)
	return ans
}

func combinationSum3Core(k, n, from int) {
	if k == 0 && n == 0 {
		ans = append(ans, slices.Clone[[]int](an))
		return
	}
	if 9-from < k {
		return
	}
	for i := from; i < 9; i++ {
		an = append(an, nums[i])
		combinationSum3Core(k-1, n-nums[i], i+1)
		an = an[:len(an)-1]
	}
}
