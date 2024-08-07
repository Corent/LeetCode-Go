package algorithm136

func singleNumber(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	return ans
}
