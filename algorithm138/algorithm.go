package algorithm138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cloneNodes(head)
	connectRandomNodes(head)
	return reconnectNodes(head)
}

func cloneNodes(head *Node) {
	ptr := head
	for ptr != nil {
		clone := &Node{Val: ptr.Val}
		clone.Next = ptr.Next
		ptr.Next = clone
		ptr = clone.Next
	}
}

func connectRandomNodes(head *Node) {
	ptr := head
	for ptr != nil {
		clone := ptr.Next
		if ptr.Random != nil {
			clone.Random = ptr.Random.Next
		}
		ptr = clone.Next
	}
}

func reconnectNodes(head *Node) *Node {
	var clonedHead, cloneNode *Node
	ptr := head
	if ptr != nil {
		cloneNode = ptr.Next
		clonedHead = cloneNode
		ptr.Next = cloneNode.Next
		ptr = ptr.Next
	}
	for ptr != nil {
		cloneNode.Next = ptr.Next
		cloneNode = cloneNode.Next
		ptr.Next = cloneNode.Next
		ptr = ptr.Next
	}
	return clonedHead
}
