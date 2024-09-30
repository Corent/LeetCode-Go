package algorithm477

func totalHammingDistance(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	cnt0s := [31]int{}
	for _, num := range nums {
		for i := 0; i < 31; i++ {
			if num&(1<<i) == 0 {
				cnt0s[i]++
			}
		}
	}
	rst, n := 0, len(nums)
	for i := 0; i < 31; i++ {
		rst += cnt0s[i] * (n - cnt0s[i])
	}
	return rst
}
