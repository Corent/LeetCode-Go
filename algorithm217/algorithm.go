package algorithm217

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}
