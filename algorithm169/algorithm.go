package algorithm169

func majorityElement(nums []int) int {
	ans, cnt := 0, 0
	for _, num := range nums {
		if cnt == 0 {
			ans = num
			cnt++
		} else if num == ans {
			cnt++
		} else {
			cnt--
		}
	}
	return ans
}
