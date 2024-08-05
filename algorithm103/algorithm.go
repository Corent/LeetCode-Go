package algorithm103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	idx, queues := 0, [][]*TreeNode{make([]*TreeNode, 0), make([]*TreeNode, 0)}
	queues[idx] = append(queues[idx], root)
	for len(queues[idx]) > 0 {
		level := make([]int, len(queues[idx]))
		for i, node := range queues[idx] {
			nth := i
			if idx == 1 {
				nth = len(queues[idx]) - i - 1
			}
			level[nth] = node.Val
			if node.Left != nil {
				queues[1-idx] = append(queues[1-idx], node.Left)
			}
			if node.Right != nil {
				queues[1-idx] = append(queues[1-idx], node.Right)
			}
		}
		ans = append(ans, level)
		queues[idx] = make([]*TreeNode, 0)
		idx = 1 - idx
	}
	return ans
}
