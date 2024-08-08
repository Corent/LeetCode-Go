package algorithm173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nums []int
	idx  int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{nums: inorderTraversal(root), idx: -1}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	lefts := inorderTraversal(root.Left)
	lefts = append(lefts, root.Val)
	return append(lefts, inorderTraversal(root.Right)...)
}

func (this *BSTIterator) Next() int {
	this.idx++
	return this.nums[this.idx]
}

func (this *BSTIterator) HasNext() bool {
	return this.idx+1 < len(this.nums)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
