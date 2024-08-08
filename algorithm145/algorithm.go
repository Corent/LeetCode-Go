package algorithm145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	ans := postorderTraversal(root.Left)
	ans = append(ans, postorderTraversal(root.Right)...)
	return append(ans, root.Val)
}
