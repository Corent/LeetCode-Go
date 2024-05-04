package algorithm004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < m && nums1[i] < nums2[j-1] { // i偏小
			iMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] { // i偏大
			iMax = i - 1
		} else { // i刚好
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else if maxLeft = nums1[i-1]; nums2[j-1] > maxLeft {
				maxLeft = nums2[j-1]
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else if minRight = nums1[i]; nums2[j] < minRight {
				minRight = nums2[j]
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}
