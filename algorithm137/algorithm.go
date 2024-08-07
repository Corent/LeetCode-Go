package algorithm137

func singleNumber(nums []int) int {
	return singleNumberK(nums, 3)
}

func singleNumberK(nums []int, k int) int {
	bits, ans := [64]int{}, 0
	for i := 0; i < 64; i++ {
		for _, n := range nums {
			bits[i] += (n >> i) & 1
		}
		ans |= (bits[i] % k) << i
	}
	return ans
}
