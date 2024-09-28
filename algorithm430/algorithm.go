package algorithm430

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	head, _ := flattenCore(root)
	return head
}

func flattenCore(root *Node) (*Node, *Node) {
	head, tail := root, root
	for tail != nil && (tail.Next != nil || tail.Child != nil) {
		next := tail.Next
		if tail.Child != nil {
			if subHead, subTail := flattenCore(tail.Child); subHead != nil && subTail != nil {
				tail.Child = nil
				tail.Next, subHead.Prev = subHead, tail
				tail = subTail
				tail.Next = next
				if next != nil {
					next.Prev = tail
				}
			}
		}
		if tail.Next != nil {
			tail = tail.Next
		}
	}
	return head, tail
}
