package algorithm129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func sumNumbers(root *TreeNode) int {
	ans = 0
	sumNumbersCore(root, 0)
	return ans
}

func sumNumbersCore(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	nextSum := sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		ans += nextSum
		return
	}
	sumNumbersCore(root.Left, nextSum)
	sumNumbersCore(root.Right, nextSum)
}
