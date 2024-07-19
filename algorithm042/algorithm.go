package algorithm042

func trap(height []int) int {
	ans := 0
	if len(height) < 3 {
		return ans
	}
	left, right := 0, len(height)-1
	for left < right {
		lowPeak := height[left]
		if lowPeak > height[right] {
			lowPeak = height[right]
		}
		if height[left] < height[right] {
			for left < right && height[left] <= lowPeak {
				ans += lowPeak - height[left]
				left++
			}
			continue
		}
		for left < right && height[right] <= lowPeak {
			ans += lowPeak - height[right]
			right--
		}
	}
	return ans
}
