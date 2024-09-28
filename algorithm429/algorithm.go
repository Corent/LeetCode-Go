package algorithm429

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	rst, idx := make([][]int, 0), 0
	queues := make([][]*Node, 2)
	queues[0] = make([]*Node, 0)
	queues[1] = make([]*Node, 0)
	queues[idx] = append(queues[idx], root)
	for len(queues[idx]) > 0 {
		vals := make([]int, 0)
		for _, node := range queues[idx] {
			if node == nil {
				continue
			}
			vals = append(vals, node.Val)
			if len(node.Children) > 0 {
				queues[1-idx] = append(queues[1-idx], node.Children...)
			}
		}
		if len(vals) > 0 {
			rst = append(rst, vals)
		}
		queues[idx] = make([]*Node, 0)
		idx = 1 - idx
	}
	return rst
}
