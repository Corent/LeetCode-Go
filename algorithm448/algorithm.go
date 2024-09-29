package algorithm448

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		idx := absInt(num) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	rst := make([]int, 0)
	for i, num := range nums {
		if num > 0 {
			rst = append(rst, i+1)
		}
	}
	return rst
}

func absInt(m int) int {
	if m < 0 {
		m = -m
	}
	return m
}
