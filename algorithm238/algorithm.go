package algorithm238

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans[i] = nums[i]
			continue
		}
		ans[i] = ans[i-1] * nums[i]
	}
	postProduct := 1
	for i := len(ans) - 1; i >= 0; i-- {
		if i == 0 {
			ans[i] = postProduct
			continue
		}
		ans[i] = ans[i-1] * postProduct
		postProduct *= nums[i]
	}
	return ans
}
