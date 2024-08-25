package algorithm380

import "math/rand"

type RandomizedSet struct {
	nums   []int
	valIdx map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.valIdx[val]; ok {
		return false
	}
	rs.valIdx[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.valIdx[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	rs.nums[id] = rs.nums[last]
	rs.valIdx[rs.nums[id]] = id
	rs.nums = rs.nums[:last]
	delete(rs.valIdx, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
