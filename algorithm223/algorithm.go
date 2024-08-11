package algorithm223

// int computeArea(int A, int B, int C, int D, int E, int F, int G, int H) {
// A: ax1
// B: ay1
// C: ax2
// D: ay2
// E: bx1
// F: by1
// G: bx2
// H: by2
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	ans := absInt(ax1-ax2)*absInt(ay1-ay2) + absInt(bx1-bx2)*absInt(by1-by2)
	if !(ay1 > by2 || ay2 < by1) && !(ax2 < bx1 || ax1 > bx2) {
		h := minInt(ay2, by2) - maxInt(ay1, by1)
		w := minInt(ax2, bx2) - maxInt(ax1, bx1)
		ans -= h * w
	}
	return ans
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func maxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}
