package algorithm104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepthChild := maxDepth(root.Left)
	if tmp := maxDepth(root.Right); tmp > maxDepthChild {
		maxDepthChild = tmp
	}
	return maxDepthChild + 1
}
