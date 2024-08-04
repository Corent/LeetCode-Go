package algorithm098

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	pre *TreeNode
)

func isValidBST(root *TreeNode) bool {
	pre = nil
	return isValidBSTCore(root)
}

func isValidBSTCore(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isValidBSTCore(root.Left) {
		if pre == nil || root.Val > pre.Val {
			pre = root
			return isValidBSTCore(root.Right)
		}
		return false
	}
	return false
}
