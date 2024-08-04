package algorithm095

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return generateTreesCore(1, n)
}

func generateTreesCore(from, to int) []*TreeNode {
	ans := make([]*TreeNode, 0)
	if from > to {
		ans = append(ans, nil)
		return ans
	}
	for root := from; root <= to; root++ {
		lefts, rights := generateTreesCore(from, root-1), generateTreesCore(root+1, to)
		for _, left := range lefts {
			for _, right := range rights {
				ans = append(ans, &TreeNode{Val: root, Left: left, Right: right})
			}
		}
	}
	return ans
}
