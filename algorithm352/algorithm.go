package algorithm352

import "slices"

type SummaryRanges struct {
	intervals [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{intervals: make([][]int, 0)}
}

func (this *SummaryRanges) AddNum(value int) {
	if len(this.intervals) == 0 {
		this.intervals = append(this.intervals, []int{value, value})
		return
	}
	if len(this.intervals) == 1 {
		if this.intervals[0][1] < value {
			p := this.intervals[0]
			if p[1] == value-1 {
				this.intervals[0][1] = value
				return
			}
			this.intervals = append(this.intervals, []int{value, value})
			return
		}
		if this.intervals[0][0] > value {
			p := this.intervals[0]
			if p[0] == value+1 {
				this.intervals[0][0] = value
				return
			}
			this.intervals = [][]int{{value, value}, p}
			return
		}
		return
	}
	pos := this.findIdx(value)
	if pos == -1 {
		if value == this.intervals[0][0]-1 {
			this.intervals[0][0] = value
		} else {
			this.intervals = append([][]int{{value, value}}, this.intervals...)
		}
		return
	}
	if pos == len(this.intervals)-1 {
		if value > this.intervals[pos][1] {
			if value == this.intervals[pos][1]+1 {
				this.intervals[pos][1] = value
			} else {
				this.intervals = append(this.intervals, []int{value, value})
			}
		}
		return
	}
	if value > this.intervals[pos][1] {
		if value == this.intervals[pos+1][0]-1 && value == this.intervals[pos][1]+1 {
			p1, p2, part2 := this.intervals[pos], this.intervals[pos+1], this.intervals[pos+2:]
			this.intervals = append(this.intervals[:pos], []int{p1[0], p2[1]})
			this.intervals = append(this.intervals, part2...)
		} else if value == this.intervals[pos+1][0]-1 {
			this.intervals[pos+1][0] = value
		} else if value == this.intervals[pos][1]+1 {
			this.intervals[pos][1] = value
		} else {
			part2 := slices.Clone[[][]int](this.intervals[pos+1:])
			this.intervals = append(this.intervals[:pos+1], []int{value, value})
			this.intervals = append(this.intervals, part2...)
		}
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.intervals
}

func (this *SummaryRanges) findIdx(val int) int {
	start, end := 0, len(this.intervals)-1
	if this.intervals[start][0] > val {
		return -1
	}
	if this.intervals[end][0] < val {
		return end
	}
	for start < end {
		mid := (start + end) >> 1
		if val < this.intervals[mid][0] {
			end = mid - 1
		} else if val > this.intervals[mid][0] {
			if val < this.intervals[mid+1][0] {
				return mid
			}
			start = mid + 1
		} else {
			return mid
		}
	}
	return start
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
