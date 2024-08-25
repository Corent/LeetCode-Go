package algorithm396

func maxRotateFunction(nums []int) int {
	f, sum, n := 0, 0, len(nums)
	for i, m := range nums {
		sum, f = sum+m, f+i*m
	}
	rst := f
	for i := 1; i <= n; i++ {
		if f = f + sum - n*nums[n-i]; f > rst {
			rst = f
		}
	}
	return rst
}
