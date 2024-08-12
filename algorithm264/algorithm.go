package algorithm264

func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	for idx, idx2, idx3, idx5 := 1, 0, 0, 0; idx < n; {
		nums[idx] = minInt(minInt(nums[idx2]*2, nums[idx3]*3), nums[idx5]*5)
		for nums[idx2]*2 <= nums[idx] {
			idx2++
		}
		for nums[idx3]*3 <= nums[idx] {
			idx3++
		}
		for nums[idx5]*5 <= nums[idx] {
			idx5++
		}
		idx++
	}
	return nums[n-1]
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}
