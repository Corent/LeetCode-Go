package algorithm1184

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start == destination {
		return 0
	}
	if destination < start {
		start, destination = destination, start
	}
	sums := make([]int, len(distance))
	for i := range distance {
		if sums[i] = distance[i]; i > 0 {
			sums[i] += sums[i-1]
		}
	}
	pathLen1 := sums[destination-1] - sums[start] + distance[start]
	if pathLen2 := sums[len(sums)-1] - pathLen1; pathLen2 < pathLen1 {
		return pathLen2
	}
	return pathLen1
}
