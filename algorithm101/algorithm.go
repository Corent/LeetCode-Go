package algorithm101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil {
		return root2 == nil
	}
	if root2 == nil || root1.Val != root2.Val {
		return false
	}
	return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}
