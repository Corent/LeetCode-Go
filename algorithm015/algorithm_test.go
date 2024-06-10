package algorithm015

import "testing"

func TestThreeSum(t *testing.T) {
	nums := []int{1, -1, -1, 0}
	ans := threeSum(nums)
	t.Log(ans)
}
