package algorithm421

func findMaximumXOR(nums []int) int {
	max, mask := 0, 0
	for i := 31; i >= 0; i-- {
		mask = mask | (1 << i)
		set := make(map[int]struct{})
		for _, num := range nums {
			set[mask&num] = struct{}{}
		}
		tmp := max | (1 << i)
		for pre := range set {
			if _, ok := set[tmp^pre]; ok {
				max = tmp
				break
			}
		}
	}
	return max
}
