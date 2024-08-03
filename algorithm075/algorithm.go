package algorithm075

func sortColors(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	idx0, idx1, idx2 := -1, n-1, n
	for idx1 > idx0 {
		switch nums[idx1] {
		case 0:
			idx0++
			nums[idx0], nums[idx1] = nums[idx1], nums[idx0]
		case 1:
			idx1--
		case 2:
			idx2--
			nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
			idx1--
		}
	}
}
