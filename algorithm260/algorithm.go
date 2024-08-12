package algorithm260

func singleNumber(nums []int) []int {
	rst, num1, num2 := 0, 0, 0
	for _, num := range nums {
		rst ^= num
	}
	rst ^= rst & (rst - 1) // 保留最高位 1，其他位置全部置零
	for _, num := range nums {
		if rst&num != 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
