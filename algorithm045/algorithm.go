package algorithm045

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	farest, prevFarest, jumpCnt := 0, 0, 0
	for i, _ := range nums {
		if prevFarest >= len(nums)-1 {
			break
		}
		for i <= prevFarest {
			if i+nums[i] > farest {
				farest = i + nums[i]
			}
			i++
		}
		prevFarest, jumpCnt = farest, jumpCnt+1
	}
	return jumpCnt
}
