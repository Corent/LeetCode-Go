package algorithm496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i, num := range nums2 {
		m[num] = i
	}
	rst := make([]int, len(nums1))
	for i, num := range nums1 {
		rst[i] = -1
		if idx, ok := m[num]; ok {
			for j := idx + 1; j < len(nums2); j++ {
				if nums2[j] > num {
					rst[i] = nums2[j]
					break
				}
			}
		}
	}
	return rst
}
