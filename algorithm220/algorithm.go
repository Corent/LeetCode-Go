package algorithm220

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, valueDiff+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= valueDiff {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= valueDiff {
			return true
		}
		mp[id] = x
		if i >= indexDiff {
			delete(mp, getID(nums[i-indexDiff], valueDiff+1))
		}
	}
	return false
}

// 若 abs(nums[i], nums[j]) < valueDiff，则 getID((nums[i]) == getID((nums[j]) 成立
// 若不存在ID相同的两个数，则比较相邻id的数
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 超时
/*func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if len(nums) == 0 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		j := 0
		if i-indexDiff > 0 {
			j = i - indexDiff
		}
		for ; j >= 0 && j < i; j++ {
			if abs(nums[i]-nums[j]) <= valueDiff {
				return true
			}
		}
	}
	return false
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}*/
