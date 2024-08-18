package algorithm324

import "fmt"

func wiggleSort(nums []int) {
	n := len(nums)
	kth := findKthSmall(nums, 0, n-1, (n+1)/2) // 后半段的第一个元素
	fmt.Println(nums, kth)
	for i, j, m := 0, 0, n-1; i <= m; {
		if a, b, c := mapIdx(i, n), mapIdx(j, n), mapIdx(m, n); nums[a] > kth {
			nums[a], nums[b] = nums[b], nums[a]
			i, j = i+1, j+1
		} else if nums[a] < kth {
			nums[a], nums[c] = nums[c], nums[a]
			m--
		} else {
			i++
		}
	}
}

func mapIdx(idx, n int) int {
	if idx*2+1 < n {
		return idx*2 + 1
	}
	return (idx - n/2) * 2
}

func findKthSmall(nums []int, from, to, k int) int {
	for from < to {
		i, j := from, to
		for i < j {
			for i < j && nums[i] <= nums[j] {
				j--
			}
			for i < j && nums[i] <= nums[j] {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		if i == k {
			return nums[i]
		}
		if i < k {
			from = i + 1
		} else {
			to = i - 1
		}
	}
	return nums[from]
}
