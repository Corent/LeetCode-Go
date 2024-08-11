package algorithm239

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || n < k || k <= 0 {
		return []int{0}
	}
	idx, ans, queue := 0, make([]int, n-k+1), make([]int, 0)
	for i := 0; i < n; i++ {
		if len(queue) > 0 && i-queue[0] == k {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			ans[idx] = nums[queue[0]]
			idx++
		}
	}
	return ans
}
