package algorithm283

func moveZeroes(nums []int) {
	idx := 0
	for _, n := range nums {
		if n == 0 {
			continue
		}
		nums[idx], idx = n, idx+1
	}
	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}
