package algorithm088

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
	}
	k := m + n
	for m > 0 && n > 0 {
		k--
		if nums1[m-1] >= nums2[n-1] {
			m--
			nums1[k] = nums1[m]
		} else {
			n--
			nums1[k] = nums2[n]
		}
	}
	for n > 0 {
		k--
		n--
		nums1[k] = nums2[n]
	}
}
