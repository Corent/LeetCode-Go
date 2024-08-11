package algorithm228

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return make([]string, 0)
	}
	ans, start := make([]string, 0), 0
	for i, n := 1, len(nums); i <= n; i++ {
		if i == n || nums[i] != nums[i-1]+1 {
			an := strconv.Itoa(nums[start])
			if start != i-1 {
				an = fmt.Sprintf("%s->%d", an, nums[i-1])
			}
			ans, start = append(ans, an), i
		}
	}
	return ans
}
