package algorithm321

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n, rst := len(nums1), len(nums2), make([]int, k)
	for i := maxInt(0, k-n); i <= minInt(k, m); i++ {
		next := merge(maxNums(nums1, i), maxNums(nums2, k-i))
		if !compare(rst, next, 0, 0) {
			rst = next
		}
	}
	return rst
}

func compare(nums1, nums2 []int, idx1, idx2 int) bool {
	for m, n := len(nums1), len(nums2); idx1 < m && idx2 < n; idx1, idx2 = idx1+1, idx2+1 {
		if nums1[idx1] > nums2[idx2] {
			return true
		}
		if nums1[idx1] < nums2[idx2] {
			return false
		}
	}
	return idx1 != len(nums1)
}

func merge(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	rst := make([]int, m+n)
	for i, j, k := 0, 0, 0; i < m || j < n; {
		if compare(nums1, nums2, i, j) {
			rst[k] = nums1[i]
			i++
		} else {
			rst[k] = nums2[j]
			j++
		}
		k++
	}
	return rst
}

func maxNums(nums []int, k int) []int {
	drop, rst := len(nums)-k, make([]int, 0)
	for _, num := range nums {
		for drop > 0 && len(rst) > 0 && rst[len(rst)-1] < num {
			rst, drop = rst[:len(rst)-1], drop-1
		}
		rst = append(rst, num)
	}
	return rst[:k]
}

func maxInt(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func minInt(m, n int) int {
	if m > n {
		return n
	}
	return m
}
