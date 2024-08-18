package algorithm324

import "testing"

func TestWiggleSort(t *testing.T) {
	nums := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums)
	t.Log(nums)
}
