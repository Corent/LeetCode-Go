package algorithm268

func missingNumber(nums []int) int {
	ans := nums[0] ^ 0
	for i := 1; i <= len(nums); i++ {
		ans ^= i
		if i != len(nums) {
			ans ^= nums[i]
		}
	}
	return ans
}
