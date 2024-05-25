package algorithm011

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		minHeight := height[i]
		if height[i] > height[j] {
			minHeight = height[j]
		}
		if v := (j - i) * minHeight; v > max {
			max = v
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
