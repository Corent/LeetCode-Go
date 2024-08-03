package algorithm085

func maximalRectangle(matrix [][]byte) int {
	heights, maxArea := make([]int, len(matrix[0])), 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if heights[j] == 0 || matrix[i][j] == '0' {
				heights[j] = int(matrix[i][j] - '0')
			} else {
				heights[j] += int(matrix[i][j] - '0')
			}
		}
		tmpMaxArea := largestRectangleArea(heights)
		if tmpMaxArea > maxArea {
			maxArea = tmpMaxArea
		}
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	maxArea, stack := 0, make([]int, 0)
	for i, n := 0, len(heights); i <= n; i++ {
		curt := -1
		if i < n {
			curt = heights[i]
		}
		for len(stack) > 0 && curt <= heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h, from := heights[idx], -1
			if len(stack) > 0 {
				from = stack[len(stack)-1]
			}
			w := i - from - 1
			area := h * w
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}
