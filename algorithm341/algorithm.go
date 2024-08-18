package algorithm341

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	Value int
	List  []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.List == nil
}

func (this NestedInteger) GetInteger() int {
	return this.Value
}

func (this *NestedInteger) SetInteger(value int) {
	this.Value = value
}

func (this *NestedInteger) Add(elem NestedInteger) {
	this.List = append(this.List, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.List
}

type NestedIterator struct {
	Stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := make([]*NestedInteger, 0)
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	return &NestedIterator{Stack: stack}
}

func (this *NestedIterator) Next() int {
	if len(this.Stack) == 0 {
		return 0
	}
	peek := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	return peek.Value
}

func (this *NestedIterator) HasNext() bool {
	for len(this.Stack) > 0 {
		peek := this.Stack[len(this.Stack)-1]
		if peek.IsInteger() {
			return true
		}
		this.Stack = this.Stack[:len(this.Stack)-1]
		list := peek.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			this.Stack = append(this.Stack, list[i])
		}
	}
	return false
}
