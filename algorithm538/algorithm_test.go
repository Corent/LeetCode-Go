package algorithm538

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	convertBSTCore(root)
	return root
}

func convertBSTCore(root *TreeNode) {
	if root == nil {
		return
	}
	convertBSTCore(root.Right)
	sum += root.Val
	root.Val = sum
	convertBSTCore(root.Left)
}
