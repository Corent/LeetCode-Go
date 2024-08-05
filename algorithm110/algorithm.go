package algorithm110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := isBalancedCore(root)
	return balanced
}

func isBalancedCore(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftBalanced := isBalancedCore(root.Left)
	if !leftBalanced {
		return 0, false
	}
	rightDepth, rightBalanced := isBalancedCore(root.Right)
	if !rightBalanced {
		return 0, false
	}
	if diff := leftDepth - rightDepth; diff < -1 || diff > 1 {
		return 0, false
	}
	maxDepthChild := leftDepth
	if maxDepthChild < rightDepth {
		maxDepthChild = rightDepth
	}
	return maxDepthChild, true
}
