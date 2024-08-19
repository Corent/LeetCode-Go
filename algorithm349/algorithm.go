package algorithm349

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = false
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			m[v] = true
		}
	}
	rst := make([]int, 0)
	for k, v := range m {
		if v {
			rst = append(rst, k)
		}
	}
	return rst
}
