package algorithm124

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	_, ans := maxPathSumCore(root)
	return ans
}

// returns: 根节点为终点的最大路径长，最大路径长(不一定经过根节点)
func maxPathSumCore(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}
	if root.Left != nil && root.Right == nil {
		maxPathEndWithLeft, maxPathInLeft := maxPathSumCore(root.Left)
		maxPathEndWithRoot := maxInt([]int{root.Val, maxPathEndWithLeft + root.Val})
		maxPath := maxInt([]int{maxPathEndWithRoot, maxPathInLeft})
		return maxPathEndWithRoot, maxPath
	}
	if root.Left == nil && root.Right != nil {
		maxPathEndWithRight, maxPathInRight := maxPathSumCore(root.Right)
		maxPathEndWithRoot := maxInt([]int{root.Val, maxPathEndWithRight + root.Val})
		maxPath := maxInt([]int{maxPathEndWithRoot, maxPathInRight})
		return maxPathEndWithRoot, maxPath
	}
	maxPathEndWithLeft, maxPathInLeft := maxPathSumCore(root.Left)
	maxPathEndWithRight, maxPathInRight := maxPathSumCore(root.Right)
	maxPathEndWithRoot := maxInt([]int{maxPathEndWithLeft, maxPathEndWithRight, 0}) + root.Val
	maxPath := maxInt([]int{maxPathEndWithRoot, maxPathEndWithLeft + maxPathEndWithRight + root.Val, maxPathInLeft, maxPathInRight})
	return maxPathEndWithRoot, maxPath
}

func maxInt(nums []int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}
