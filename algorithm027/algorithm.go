package algorithm027

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, len(nums)-1
	for i < j {
		for ; i < j && nums[i] != val; i++ {
		}
		for ; i < j && nums[j] == val; j-- {
		}
		nums[i], nums[j] = nums[j], nums[i]
		i, j = i+1, j-1
	}
	if nums[i] == val {
		return i
	}
	return i + 1
}
