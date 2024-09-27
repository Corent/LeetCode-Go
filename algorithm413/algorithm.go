package algorithm413

func numberOfArithmeticSlices(nums []int) int {
	cnt, addend := 0, 0
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i] == nums[i-2]-nums[i-1] {
			addend++
			cnt += addend
		} else {
			addend = 0
		}
	}
	return cnt
}
