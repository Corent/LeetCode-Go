package algorithm116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	idx, queues := 0, [][]*Node{make([]*Node, 0), make([]*Node, 0)}
	queues[idx] = append(queues[idx], root)
	for len(queues[idx]) > 0 {
		for i, node := range queues[idx] {
			if i < len(queues[idx])-1 {
				node.Next = queues[idx][i+1]
			}
			if node.Left != nil {
				queues[1-idx] = append(queues[1-idx], node.Left)
			}
			if node.Right != nil {
				queues[1-idx] = append(queues[1-idx], node.Right)
			}
		}
		queues[idx] = make([]*Node, 0)
		idx = 1 - idx
	}
	return root
}
