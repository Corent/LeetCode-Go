package algorithm398

import "math/rand"

type Solution map[int][]int

func Constructor(nums []int) Solution {
	pos := map[int][]int{}
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}
	return pos
}

func (pos Solution) Pick(target int) int {
	indices := pos[target]
	return indices[rand.Intn(len(indices))]
}

// 超时
/*type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {
	rst, cnt := 0, 0
	for i, num := range this.nums {
		if num == target {
			if cnt++; rand.Intn(cnt) == 0 {
				rst = i
			}
		}
	}
	return rst
}*/
