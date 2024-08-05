package algorithm111

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var minDp int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDp = math.MaxInt
	dfsTravel(root, 0)
	return minDp
}

func dfsTravel(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		if depth < minDp {
			minDp = depth
		}
		return
	}
	dfsTravel(root.Left, depth)
	dfsTravel(root.Right, depth)
}
