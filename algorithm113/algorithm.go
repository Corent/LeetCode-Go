package algorithm113

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ans [][]int
	cur []int
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans, cur = make([][]int, 0), make([]int, 0)
	dfsTravel(root, targetSum)
	return ans
}

func dfsTravel(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		ans = append(ans, slices.Clone[[]int](cur))
	}
	dfsTravel(root.Left, targetSum-root.Val)
	dfsTravel(root.Right, targetSum-root.Val)
	cur = cur[:len(cur)-1]
}
