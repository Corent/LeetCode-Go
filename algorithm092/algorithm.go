package algorithm092

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var pre, p *ListNode // p为反转前left的前序节点
	for idx, cur := 1, head; cur != nil; {
		if idx == right { // 4. 遍历到了right
			if p != nil { // left != 1 的情况，p存在
				p.Next.Next = cur.Next
				cur.Next = pre
				p.Next = cur
			} else { // left == 1 情况，p不存在
				head.Next = cur.Next
				head = cur
				cur.Next = pre
			}
			break
		}
		if idx < left { // 1. 还未遍历超过left
			pre = cur
			cur = cur.Next
		} else if idx == left { // 2. 已经遍历超过left，记录left前序节点
			p = pre
			pre = cur
			cur = cur.Next
		} else { // 3. pre、cur 位置分别 left、left + 1，开始反转
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}
		idx++
	}
	return head
}
