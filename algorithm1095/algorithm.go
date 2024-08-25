package algorithm1095

type MountainArray struct {
	A      []int
	GetCnt int
}

func (m *MountainArray) get(idx int) int {
	defer func() {
		m.GetCnt++
	}()
	if m.GetCnt > 100 {
		panic("ExceededMaximumCalledTimes")
	}
	return m.A[idx]
}

func (m *MountainArray) length() int {
	return len(m.A)
}

func Construct(a []int) *MountainArray {
	return &MountainArray{A: a}
}

func FindInMountainArray(target int, mountainArr *MountainArray) int {
	peek := findPeak(mountainArr)
	if target == mountainArr.get(peek) {
		return peek
	}
	if leftIdx := binarySearch(mountainArr, 0, peek-1, target, true); leftIdx != -1 {
		return leftIdx
	}
	return binarySearch(mountainArr, peek+1, mountainArr.length()-1, target, false)
}

func findPeak(mountainArr *MountainArray) int {
	left, right := 0, mountainArr.length()-1
	for left+1 < right {
		mid := (left + right) >> 1
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			left = mid
		} else {
			right = mid
		}
	}
	if mountainArr.get(left) > mountainArr.get(right) {
		return left
	}
	return right
}

func binarySearch(mountainArr *MountainArray, left, right, target int, flag bool) int {
	for left+1 < right {
		mid := (left + right) >> 1
		if midVal := mountainArr.get(mid); midVal == target {
			return mid
		} else if midVal > target {
			if flag {
				right = mid
			} else {
				left = mid
			}
		} else {
			if flag {
				left = mid
			} else {
				right = mid
			}
		}
	}
	if mountainArr.get(left) == target {
		return left
	}
	if mountainArr.get(right) == target {
		return right
	}
	return -1
}
