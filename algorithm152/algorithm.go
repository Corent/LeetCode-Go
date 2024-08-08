package algorithm152

// https://www.cnblogs.com/gyyyl/p/14397411.html

func maxProduct(nums []int) int {
	maxs, mins, ans := make([]int, len(nums)), make([]int, len(nums)), nums[0]
	maxs[0], mins[0] = nums[0], nums[0]
	for index := 1; index < len(nums); index++ {
		if nums[index] == 0 {
			maxs[index], mins[index] = 0, 0
		} else {
			var tmpMax, tmpMin int
			if nums[index] > 0 {
				tmpMax = nums[index] * maxs[index-1]
				tmpMin = nums[index] * mins[index-1]
			} else {
				tmpMax = nums[index] * mins[index-1]
				tmpMin = nums[index] * maxs[index-1]
			}
			if tmpMax > nums[index] {
				maxs[index] = tmpMax
			} else {
				maxs[index] = nums[index]
			}
			if tmpMin < nums[index] {
				mins[index] = tmpMin
			} else {
				mins[index] = nums[index]
			}
		}
		if ans < maxs[index] {
			ans = maxs[index]
		}
	}
	return ans
}

// 最小值溢出 for case []int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10, -10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0}
/*func maxProduct(nums []int) int {
	ans, maxVal, minVal := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxTmp, minTmp := maxVal*nums[i], minVal*nums[i]
		maxVal = maxInt(maxInt(maxTmp, minTmp), nums[i])
		minVal = minInt(minInt(maxTmp, minTmp), nums[i])
		ans = maxInt(ans, maxVal)
		fmt.Println("max val:", maxVal, "min val:", minVal, "ans:", ans)
	}
	return ans
}

func minInt(m, n int) int {
	minVal, _ := minMax(m, n)
	return minVal
}

func maxInt(m, n int) int {
	_, maxVal := minMax(m, n)
	return maxVal
}

func minMax(m, n int) (int, int) {
	if m > n {
		return n, m
	}
	return m, n
}*/
