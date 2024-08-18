package algorithm327

import "slices"

// https://www.bilibili.com/video/BV1Zo4y1j75u
var (
	sums         []int
	Lower, Upper int
)

func countRangeSum(nums []int, lower int, upper int) int {
	sums, Lower, Upper = []int{0}, lower, upper
	for _, num := range nums {
		sums = append(sums, sums[len(sums)-1]+num)
	}
	return merge(0, len(nums))
}

func merge(left, right int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) >> 1
	cnt := merge(left, mid) + merge(mid+1, right)
	for i, j, k := left, mid, mid; i <= mid && k <= right; i++ {
		for j+1 <= right && sums[j+1] < sums[i]+Lower {
			j++
		}
		for k+1 <= right && sums[k+1] <= sums[i]+Upper {
			k++
		}
		cnt += k - j
	}
	tmp := slices.Clone[[]int](sums[left : mid+1])
	for i, j, k := 0, mid+1, left; i <= mid-left || j <= right; {
		if i <= mid-left && j > right {
			sums[k] = tmp[i]
			i, k = i+1, k+1
		} else if i > mid-left && j <= right {
			sums[k] = sums[j]
			j, k = j+1, k+1
		} else if tmp[i] < sums[j] {
			sums[k] = tmp[i]
			i, k = i+1, k+1
		} else {
			sums[k] = sums[j]
			j, k = j+1, k+1
		}
	}
	return cnt
}

// 超时
/*var (
	NUMS, sums   []int
	Lower, Upper int
)

func countRangeSum(nums []int, lower int, upper int) int {
	NUMS, Lower, Upper = nums, lower, upper
	n := len(nums)
	if n == 0 {
		return 0
	}
	sums = make([]int, n)
	sums[0] = nums[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	return countNum(0, n-1)
}

func countNum(left, right int) int {
	if left == right {
		if NUMS[left] >= Lower && NUMS[right] <= Upper {
			return 1
		}
		return 0
	}
	mid, cnt := (left+right)>>1, 0
	for i := left; i <= mid; i++ {
		for j := mid + 1; j <= right; j++ {
			if sum := sums[j] - sums[i] + NUMS[i]; sum >= Lower && sum <= Upper {
				cnt++
			}
		}
	}
	return cnt + countNum(left, mid) + countNum(mid+1, right)
}
*/
