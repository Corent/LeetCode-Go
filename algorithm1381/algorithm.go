package algorithm1381

type CustomStack struct {
	Stack        []int
	Idx, MaxSize int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		Stack:   make([]int, maxSize),
		MaxSize: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if this.Idx < this.MaxSize {
		if this.Idx < 0 {
			this.Idx = 0
		}
		this.Stack[this.Idx] = x
		this.Idx++
	}
}

func (this *CustomStack) Pop() int {
	tmp := -1
	if this.Idx > 0 {
		tmp = this.Stack[this.Idx-1]
	}
	this.Idx--
	return tmp
}

func (this *CustomStack) Increment(k int, val int) {
	if k > this.Idx {
		k = this.Idx
	}
	for i := 0; i < k; i++ {
		this.Stack[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
