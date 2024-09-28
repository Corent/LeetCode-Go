package algorithm437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	rst := pathSumRoot(root, targetSum)
	rst += pathSum(root.Left, targetSum)
	rst += pathSum(root.Right, targetSum)
	return rst
}

func pathSumRoot(root *TreeNode, targetSum int) (rst int) {
	if root == nil {
		return
	}
	if root.Val == targetSum {
		rst++
	}
	if root.Left != nil {
		rst += pathSumRoot(root.Left, targetSum-root.Val)
	}
	if root.Right != nil {
		rst += pathSumRoot(root.Right, targetSum-root.Val)
	}
	return rst
}
