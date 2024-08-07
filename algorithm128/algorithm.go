package algorithm128

func longestConsecutive(nums []int) int {
	ans := 0
	if len(nums) == 0 {
		return ans
	}

	set := make(map[int]bool, len(nums))
	for _, v := range nums {
		set[v] = true
	}
	for _, n := range nums {
		/*if !set[n] { // 低效
			continue
		}*/
		prev, next := n-1, n+1
		for set[prev] {
			delete(set, prev)
			prev--
		}
		for set[next] {
			delete(set, next)
			next++
		}
		if length := next - prev - 1; length > ans {
			ans = length
		}
	}
	return ans
}
