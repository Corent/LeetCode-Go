package algorithm313

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	nums, idxes := make([]int, n), make([]int, len(primes))
	nums[0] = 1
	for idx := 1; idx < n; {
		nums[idx] = math.MaxInt
		for i := range idxes {
			nums[idx] = minInt(nums[idx], nums[idxes[i]]*primes[i])
		}
		for i := range idxes {
			for nums[idxes[i]]*primes[i] <= nums[idx] {
				idxes[i]++
			}
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
