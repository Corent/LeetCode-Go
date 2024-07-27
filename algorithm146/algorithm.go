package algorithm146

type LRUCache struct {
	Capacity, Size int
	Head, Tail     *Node         // head -> tail 上次访问时间 远 -> 近
	Prevs          map[int]*Node // 保存上一个节点
}

type Node struct {
	Key, Value int
	Next       *Node
}

func Constructor(capacity int) LRUCache {
	node := &Node{}
	return LRUCache{Capacity: capacity, Head: node, Tail: node, Prevs: make(map[int]*Node, capacity)}
}

func (this *LRUCache) Get(key int) int {
	prev, ok := this.Prevs[key]
	if !ok {
		return -1
	}
	this.moveToTail(prev)
	return this.Prevs[key].Next.Value
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.Prevs[key]
	if ok {
		this.Prevs[key].Next.Value = value
		this.moveToTail(this.Prevs[key])
		return
	}
	node := &Node{Key: key, Value: value}
	this.Tail.Next = node
	this.Prevs[key] = this.Tail
	this.Tail = node
	this.Size++
	if this.Size > this.Capacity {
		delete(this.Prevs, this.Head.Next.Key)
		this.Head.Next = this.Head.Next.Next
		if this.Head.Next != nil {
			this.Prevs[this.Head.Next.Key] = this.Head
		}
		this.Size--
	}
}

func (this *LRUCache) moveToTail(prev *Node) {
	if prev.Next == this.Tail {
		return
	}
	node := prev.Next
	prev.Next = node.Next
	if node.Next != nil {
		this.Prevs[node.Next.Key] = prev
	}
	this.Tail.Next = node
	node.Next = nil
	this.Prevs[node.Key] = this.Tail
	this.Tail = node
}
