package algorithm996

import (
	"math"
	"sort"
)

func numSquarefulPerms(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	st := make([]bool, n)

	var dfs func(start, p int) int
	dfs = func(start, p int) int {
		if start == n {
			return 1
		}

		var s int
		for i := 0; i < n; i++ {
			// 全排列的剪枝思路
			if i > 0 && nums[i] == nums[i-1] && !st[i-1] {
				continue
			}
			if p == -1 || !st[i] && check(nums[p]+nums[i]) {
				st[i] = true
				s += dfs(start+1, i)
				st[i] = false
			}
		}
		return s
	}

	return dfs(0, -1)
}

func check(x int) bool {
	t := int(math.Sqrt(float64(x)))
	return t*t == x
}
