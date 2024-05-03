package algorithm001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if i == 0 {
			m[v] = i
			continue
		}
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
}
