package algorithm164

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxVal, minVal := nums[0], nums[0]
	for _, num := range nums[1:] {
		maxVal = maxInt(maxVal, num)
		minVal = minInt(minVal, num)
	}
	length := (maxVal-minVal)/len(nums) + 1
	buckets := make([][]int, (maxVal-minVal)/length+1)
	for _, num := range nums {
		i := (num - minVal) / length
		if buckets[i] == nil {
			buckets[i] = []int{num, num}
		} else {
			buckets[i][0] = minInt(buckets[i][0], num)
			buckets[i][1] = maxInt(buckets[i][1], num)
		}
	}

	gap, prev := 0, 0
	for i, bucket := range buckets {
		if bucket == nil {
			continue
		}
		gap = maxInt(gap, bucket[0]-buckets[prev][1])
		prev = i
	}
	return gap
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func maxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}
