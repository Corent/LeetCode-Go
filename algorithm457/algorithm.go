package algorithm457

var (
	Nums []int
	n    int
)

func circularArrayLoop(nums []int) bool {
	if Nums, n = nums, len(nums); n < 2 {
		return false
	}
	for i := 0; i < len(Nums); i++ {
		slow, fast := i, nextIdx(i)
		for Nums[i]*Nums[slow] > 0 && Nums[i]*Nums[fast] > 0 && Nums[i]*Nums[nextIdx(fast)] > 0 {
			if slow == fast {
				if slow == nextIdx(slow) {
					break
				}
				return true
			}
			slow, fast = nextIdx(slow), nextIdx(nextIdx(fast))
		}
	}
	return false
}

func nextIdx(i int) int {
	next := i + Nums[i]
	if next >= 0 {
		return next % n
	}
	return next%n + n
}
