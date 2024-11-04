package algorithm754

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	sum, idx := 0, 0
	for sum < target {
		idx++
		sum += idx
	}
	if sum == target || (sum-target)&1 == 0 {
		return idx
	}
	if idx&1 == 0 {
		return idx + 1
	}
	return idx + 2
}
