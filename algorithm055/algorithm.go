package algorithm055

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	farest, prevFarest := 0, 0
	for i := 0; i < len(nums); {
		if prevFarest >= len(nums)-1 {
			return true
		}
		for ; i <= prevFarest; i++ {
			if i+nums[i] > farest {
				farest = i + nums[i]
			}
		}
		prevFarest = farest
		if i > farest {
			break
		}
	}
	return false
}
