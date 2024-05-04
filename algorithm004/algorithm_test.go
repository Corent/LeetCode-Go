package algorithm004

import "testing"

func TestAlgorithm(t *testing.T) {
	nums1, nums2 := []int{1, 2, 3, 4, 5}, []int{3, 2, 1}
	nums1, nums2 = nums2, nums1
	t.Log(nums1, nums2)
}
