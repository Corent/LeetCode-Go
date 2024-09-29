package algorithm442

func findDuplicates(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	rst := make([]int, 0)
	for _, num := range nums {
		if idx := absInt(num) - 1; nums[idx] < 0 {
			rst = append(rst, absInt(num))
		} else {
			nums[idx] = -nums[idx]
		}
	}
	return rst
}

func absInt(m int) int {
	if m < 0 {
		return -m
	}
	return m
}
