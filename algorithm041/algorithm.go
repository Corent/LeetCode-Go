package algorithm041

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] > 0 && nums[i] < len(nums) && nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			continue
		}
		i++
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
