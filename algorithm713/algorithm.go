package algorithm713

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	m, n := 1, 0
	for i, j := 0, 0; j < len(nums); j++ {
		m *= nums[j]
		for m >= k {
			m /= nums[i]
			i++
		}
		n += j - i + 1
	}
	return n
}
