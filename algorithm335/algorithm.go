package algorithm335

func isSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		if i > 3 && distance[i-1] == distance[i-3] && distance[i-2] <= distance[i]+distance[i-4] {
			return true
		}
		if i > 4 && distance[i-1] <= distance[i-3] && distance[i-4] <= distance[i-2] && distance[i-2] <= distance[i]+distance[i-4] && distance[i-3] <= distance[i-1]+distance[i-5] {
			return true
		}
	}
	return false
}
