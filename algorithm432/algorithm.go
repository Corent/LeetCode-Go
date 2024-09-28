package algorithm432

type AllOne struct {
	KeyMap     map[string]int
	ValMap     map[int]*DListNode
	Head, Tail *DListNode
}

type DListNode struct {
	Val        int
	keySet     map[string]struct{}
	Prev, Next *DListNode
}

func Constructor() AllOne {
	rst := AllOne{
		KeyMap: make(map[string]int),
		ValMap: make(map[int]*DListNode),
		Head:   &DListNode{},
		Tail:   &DListNode{},
	}
	rst.Head.Next, rst.Tail.Prev = rst.Tail, rst.Head
	return rst
}

func (this *AllOne) insertNode(key string, curVal, dir int) {
	newNode := &DListNode{keySet: make(map[string]struct{})}
	this.ValMap[curVal+dir] = newNode
	newNode.keySet[key] = struct{}{}
	if curVal == 0 {
		newNode.Next = this.Head.Next
		this.Head.Next.Prev = newNode
		newNode.Prev = this.Head
		this.Head.Next = newNode
		return
	}
	curNode := this.ValMap[curVal]
	if dir == 1 {
		newNode.Next = curNode.Next
		curNode.Next.Prev = newNode
		newNode.Prev = curNode
		curNode.Next = newNode
	} else if dir == -1 {
		newNode.Prev = curNode.Prev
		newNode.Next = curNode
		curNode.Prev.Next = newNode
		curNode.Prev = newNode
	}
}

func (this *AllOne) removeNode(curVal int) {
	curNode := this.ValMap[curVal]
	if len(curNode.keySet) != 0 {
		return
	}
	delete(this.ValMap, curVal)
	curNode.Next.Prev = curNode.Prev
	curNode.Prev.Next = curNode.Next
	curNode.Next, curNode.Prev = nil, nil
}

func (this *AllOne) Inc(key string) {
	if curVal, ok := this.KeyMap[key]; ok {
		delete(this.ValMap[curVal].keySet, key)
		this.KeyMap[key] = curVal + 1
		if node, ok := this.ValMap[curVal+1]; ok {
			node.keySet[key] = struct{}{}
		} else {
			this.insertNode(key, curVal, 1)
		}
		this.removeNode(curVal)
	} else {
		this.KeyMap[key] = 1
		if node, ok := this.ValMap[1]; ok {
			node.keySet[key] = struct{}{}
		} else {
			this.insertNode(key, 0, 1)
		}
	}
}

func (this *AllOne) Dec(key string) {
	if curVal, ok := this.KeyMap[key]; ok {
		curNode := this.ValMap[curVal]
		delete(curNode.keySet, key)
		if curVal == 1 {
			delete(this.KeyMap, key)
			this.removeNode(curVal)
			return
		}
		this.KeyMap[key] = curVal - 1
		if node, ok := this.ValMap[curVal-1]; ok {
			node.keySet[key] = struct{}{}
		} else {
			this.insertNode(key, curVal, -1)
		}
		this.removeNode(curVal)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.Tail.Prev == this.Head {
		return ""
	}
	for k, _ := range this.Tail.Prev.keySet {
		return k
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.Head.Next == this.Tail {
		return ""
	}
	for k, _ := range this.Head.Next.keySet {
		return k
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
