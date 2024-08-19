package algorithm349

import "testing"

func TestIntersection(t *testing.T) {
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	t.Log(intersection(nums1, nums2))
}
