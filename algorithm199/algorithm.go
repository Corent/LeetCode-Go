package algorithm199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	idx, queues := 0, [][]*TreeNode{make([]*TreeNode, 0), make([]*TreeNode, 0)}
	queues[idx] = append(queues[idx], root)
	for len(queues[idx]) > 0 {
		ans = append(ans, queues[idx][len(queues[idx])-1].Val)
		for _, node := range queues[idx] {
			if node.Left != nil {
				queues[1-idx] = append(queues[1-idx], node.Left)
			}
			if node.Right != nil {
				queues[1-idx] = append(queues[1-idx], node.Right)
			}
		}
		queues[idx] = make([]*TreeNode, 0)
		idx = 1 - idx
	}
	return ans
}
