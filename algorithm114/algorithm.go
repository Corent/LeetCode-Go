package algorithm114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	left, right := root.Left, root.Right
	flatten(left)
	flatten(right)
	root.Left = nil
	if left != nil {
		root.Right = left
		ptr := root.Right
		for ; ptr.Right != nil; ptr = ptr.Right {
		}
		ptr.Right = right
	}
}
