package algorithm219

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	count := make(map[int]struct{})
	for i, v := range nums {
		if i > k {
			delete(count, nums[i-k-1])
		}
		if _, ok := count[v]; ok {
			return true
		}
		count[v] = struct{}{}
	}
	return false
}
