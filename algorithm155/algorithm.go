package algorithm155

type MinStack struct {
	numStack, minStack []int
}

func Constructor() MinStack {
	return MinStack{numStack: make([]int, 0), minStack: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.numStack = append(this.numStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, minInt(this.minStack[len(this.minStack)-1], val))
	}
}

func (this *MinStack) Pop() {
	this.numStack = this.numStack[:len(this.numStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.numStack[len(this.numStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func minInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
