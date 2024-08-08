package algorithm179

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	sort.Slice(nums, func(i, j int) bool {
		ab, _ := strconv.ParseInt(fmt.Sprintf("%d%d", nums[i], nums[j]), 0, 64)
		ba, _ := strconv.ParseInt(fmt.Sprintf("%d%d", nums[j], nums[i]), 0, 64)
		return ab > ba
	})
	idx := 0
	for nums[idx] == 0 && idx < len(nums)-1 {
		idx++
	}
	builder := strings.Builder{}
	for idx < len(nums) {
		builder.WriteString(strconv.Itoa(nums[idx]))
	}
	return builder.String()
}
