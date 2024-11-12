package algorithm995

func minKBitFlips(nums []int, k int) int {
	if len(nums) < 0 {
		return -1
	}
	window, rst := make([]int, 0), 0
	for i, num := range nums {
		for len(window) > 0 && window[0]+k <= i {
			window = window[1:]
		}
		if num == len(window)&1 {
			if i+k > len(nums) {
				return -1
			}
			window, rst = append(window, i), rst+1
		}
	}
	return rst
}
