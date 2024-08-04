package algorithm094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans []int

func inorderTraversal(root *TreeNode) []int {
	ans = make([]int, 0)
	inorderTraversalCore(root)
	return ans
}

func inorderTraversalCore(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalCore(root.Left)
	ans = append(ans, root.Val)
	inorderTraversalCore(root.Right)
}
