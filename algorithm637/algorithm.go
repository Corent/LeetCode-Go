package algorithm637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	queue, rst := []*TreeNode{root}, make([]float64, 0)
	for len(queue) > 0 {
		nextQueue, sum := make([]*TreeNode, 0), 0
		for _, node := range queue {
			sum += node.Val
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		rst = append(rst, float64(sum)/float64(len(queue)))
		queue = nextQueue
	}
	return rst
}
