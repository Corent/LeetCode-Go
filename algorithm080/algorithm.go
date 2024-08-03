package algorithm080

func removeDuplicates(nums []int) int {
	cnt := 0
	for i, numsLen := 0, len(nums); i < numsLen; i++ {
		if cnt < 2 ||
			nums[cnt-2] != nums[i] && nums[cnt-1] == nums[i] ||
			nums[cnt-1] != nums[i] {
			nums[cnt], cnt = nums[i], cnt+1
		}
	}
	return cnt
}
