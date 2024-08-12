package algorithm275

func hIndex(citations []int) int {
	left, right, n := 0, len(citations)-1, len(citations)
	for left <= right {
		mid := (left + right) >> 1
		if citations[mid] >= n-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return n - left
}
