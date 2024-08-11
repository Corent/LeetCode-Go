package algorithm229

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return make([]int, 0)
	}
	n1, cnt1, n2, cnt2 := 0, 0, 0, 0
	for _, num := range nums {
		if num == n1 {
			cnt1++
		} else if num == n2 {
			cnt2++
		} else if cnt1 == 0 {
			n1, cnt1 = num, 1
		} else if cnt2 == 0 {
			n2, cnt2 = num, 1
		} else {
			cnt1, cnt2 = cnt1-1, cnt2-1
		}
	}
	cnt1, cnt2 = 0, 0
	for _, num := range nums {
		if num == n1 {
			cnt1++
		} else if num == n2 {
			cnt2++
		}
	}
	ans, cnt := make([]int, 0), len(nums)/3
	if cnt1 > cnt {
		ans = append(ans, n1)
	}
	if cnt2 > cnt {
		ans = append(ans, n2)
	}
	return ans
}
