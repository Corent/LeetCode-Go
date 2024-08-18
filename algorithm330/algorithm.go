package algorithm330

func minPatches(nums []int, n int) int {
	ans := 0
	for i, bound, m := 0, 1, len(nums); bound <= n; {
		if i < m && nums[i] <= bound {
			bound, i = bound+nums[i], i+1
		} else {
			bound, ans = bound<<1, ans+1
		}
	}
	return ans
}
