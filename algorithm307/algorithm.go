package algorithm307

type NumArray struct {
	root *SgmTreeNode
}

type SgmTreeNode struct {
	Val, Start, End int
	Left, Right     *SgmTreeNode
}

func Constructor(nums []int) NumArray {
	return NumArray{root: buildSgmTree(nums, 0, len(nums)-1)}
}

func buildSgmTree(nums []int, start, end int) *SgmTreeNode {
	if start > end {
		return nil
	}
	root := &SgmTreeNode{Start: start, End: end}
	if start == end {
		root.Val = nums[start]
		return root
	}
	mid := (start + end) >> 1
	root.Left = buildSgmTree(nums, start, mid)
	root.Right = buildSgmTree(nums, mid+1, end)
	root.Val = root.Left.Val + root.Right.Val
	return root
}

func (this *NumArray) Update(index int, val int) {
	update(this.root, index, val)
}

func update(root *SgmTreeNode, idx, val int) {
	if root.Start == root.End && root.Start == idx {
		root.Val = val
		return
	}
	mid := (root.End + root.Start) >> 1
	if idx <= mid {
		update(root.Left, idx, val)
	} else {
		update(root.Right, idx, val)
	}
	root.Val = root.Left.Val + root.Right.Val
}

func (this *NumArray) SumRange(left int, right int) int {
	return sumRange(this.root, left, right)
}

func sumRange(root *SgmTreeNode, left, right int) int {
	if root.Start == left && root.End == right {
		return root.Val
	}
	mid := (root.End + root.Start) >> 1
	if right <= mid {
		return sumRange(root.Left, left, right)
	}
	if left > mid {
		return sumRange(root.Right, left, right)
	}
	return sumRange(root.Left, left, mid) + sumRange(root.Right, mid+1, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
