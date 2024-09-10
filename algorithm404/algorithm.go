package algorithm404

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	sumOfLeftLeavesCore(root, false)
	return sum
}

func sumOfLeftLeavesCore(root *TreeNode, isLeft bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if isLeft {
			sum += root.Val
		}
		return
	}
	if root.Right != nil {
		sumOfLeftLeavesCore(root.Right, false)
	}
	if root.Left != nil {
		sumOfLeftLeavesCore(root.Left, true)
	}
}
