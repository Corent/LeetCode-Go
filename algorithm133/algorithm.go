package algorithm133

type Node struct {
	Val       int
	Neighbors []*Node
}

var visited map[int]*Node

func cloneGraph(node *Node) *Node {
	visited = make(map[int]*Node)
	return cloneGraphCore(node)
}

func cloneGraphCore(node *Node) *Node {
	if node == nil {
		return node
	}
	if n, ok := visited[node.Val]; ok {
		return n
	}
	newNode := &Node{Val: node.Val}
	visited[newNode.Val] = newNode
	for _, n := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, cloneGraphCore(n))
	}
	return newNode
}
