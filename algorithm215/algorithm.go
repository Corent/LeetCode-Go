package algorithm215

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

// 超时，极端case nums = [1, 2, 3, 4, 5 ...... 1], k = 5000
/*func findKthLargest(nums []int, k int) int {
	if slices.IsSorted[[]int](nums) {
		return nums[len(nums)-k]
	}
	from, to := 0, len(nums)-1
	for k = len(nums) - k; from < to; {
		i, j := from, to
		for i < j {
			for i < j && nums[j] >= nums[i] {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
			for i < j && nums[i] <= nums[j] {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		if i == from {
			return nums[i]
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
}*/
