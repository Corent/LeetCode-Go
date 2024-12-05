package algorithm1282

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)
	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)
	}
	rst := make([][]int, 0)
	for size, people := range groups {
		for i := 0; i < len(people); i += size {
			rst = append(rst, people[i:i+size])
		}
	}
	return rst
}
