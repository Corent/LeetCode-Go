package algorithm315

func countSmaller(nums []int) []int {
	smaller, pos := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(pos); i++ {
		pos[i] = i
	}
	sort(nums, smaller, pos, 0, len(nums)-1)
	return smaller
}

func sort(nums, smaller, pos []int, from, to int) {
	if from >= to {
		return
	}
	m := (from + to) >> 1
	sort(nums, smaller, pos, from, m)
	sort(nums, smaller, pos, m+1, to)
	merged := make([]int, to-from+1)
	for i, j, k, jump := from, m+1, 0, 0; i <= m || j <= to; {
		if i > m {
			merged[k] = pos[j]
			k, j, jump = k+1, j+1, jump+1
		} else if j > to {
			smaller[pos[i]] += jump
			merged[k] = pos[i]
			k, i = k+1, i+1
		} else if nums[pos[i]] <= nums[pos[j]] {
			smaller[pos[i]] += jump
			merged[k] = pos[i]
			k, i = k+1, i+1
		} else {
			merged[k] = pos[j]
			k, j, jump = k+1, j+1, jump+1
		}
	}
	for p, mg := range merged {
		pos[from+p] = mg
	}
}
