package algorithm460

import (
	"fmt"
	"sort"
)

type LFUCache struct {
	Capacity, MinCnt int
	KeyVals, keyCnts map[int]int       // 保存key、value；保存key对应value的访问次数
	CntKeys          map[int]*LRUCache // 保存相同访问次数的key的集合
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Capacity: capacity,
		MinCnt:   -1,
		KeyVals:  make(map[int]int),
		keyCnts:  make(map[int]int),
		CntKeys:  make(map[int]*LRUCache),
	}
}

func (this *LFUCache) Get(key int) int {
	defer func() {
		fmt.Printf("[Get %d]\n", key)
		this.Print()
	}()
	if _, ok := this.KeyVals[key]; !ok {
		return -1
	}
	cnt := this.keyCnts[key]
	this.keyCnts[key] = cnt + 1
	this.CntKeys[cnt].Remove(key)

	if cnt == this.MinCnt && this.CntKeys[cnt].Len() == 0 {
		this.MinCnt++
	}

	lruCache, ok := this.CntKeys[cnt+1]
	if !ok || lruCache == nil {
		this.CntKeys[cnt+1] = BuildLRUCache()
	}
	this.CntKeys[cnt+1].Put(key)

	return this.KeyVals[key]
}

func (this *LFUCache) Put(key int, value int) {
	defer func() {
		fmt.Printf("[Put %d %d]\n", key, value)
		this.Print()
	}()
	if this.Capacity <= 0 {
		return
	}

	if _, ok := this.KeyVals[key]; ok {
		this.KeyVals[key] = value
		this.Get(key)
		return
	}

	if len(this.KeyVals) >= this.Capacity {
		if k := this.CntKeys[this.MinCnt].RemoveLatest(); k >= 0 {
			delete(this.KeyVals, k)
			delete(this.keyCnts, k)
		}
	}

	this.KeyVals[key], this.MinCnt = value, 1
	this.keyCnts[key] = this.MinCnt
	if _, ok := this.CntKeys[this.MinCnt]; !ok {
		this.CntKeys[this.MinCnt] = BuildLRUCache()
	}
	this.CntKeys[this.MinCnt].Put(key)
}

func (this *LFUCache) Print() {
	fmt.Printf("MinCnt: %d\n", this.MinCnt)
	fmt.Printf("KeyVals: %v\n", this.KeyVals)
	fmt.Printf("KeyCnts: %v\n", this.keyCnts)
	cnts := make([]int, 0)
	for cnt, _ := range this.CntKeys {
		cnts = append(cnts, cnt)
	}
	sort.Ints(cnts)
	for _, cnt := range cnts {
		lruCache := this.CntKeys[cnt]
		fmt.Printf("CntKeys cnt: %d, ", cnt)
		for node := lruCache.Head.Next; node != nil; node = node.Next {
			fmt.Printf("%d -> ", node.Key)
		}
		fmt.Println("nil")
	}
}

// copied from Algorithm146
type LRUCache struct {
	Head, Tail *Node         // head -> tail 上次访问时间 远 -> 近
	Prevs      map[int]*Node // 保存上一个节点
}

type Node struct {
	Key  int
	Next *Node
}

func BuildLRUCache() *LRUCache {
	node := &Node{}
	return &LRUCache{Head: node, Tail: node, Prevs: make(map[int]*Node)}
}

func (this *LRUCache) Put(key int) {
	_, ok := this.Prevs[key]
	if ok {
		this.moveToTail(this.Prevs[key])
		return
	}
	node := &Node{Key: key}
	this.Tail.Next = node
	this.Prevs[key] = this.Tail
	this.Tail = node
}

func (this *LRUCache) Remove(key int) {
	prev, ok := this.Prevs[key]
	if !ok || prev == nil || prev.Next == nil {
		return
	}
	if prev.Next == this.Tail {
		this.Tail = prev
	}
	prev.Next = prev.Next.Next
	if prev.Next != nil {
		this.Prevs[prev.Next.Key] = prev
	}
	delete(this.Prevs, key)
}

func (this *LRUCache) RemoveLatest() int {
	if this.Head.Next == nil {
		return -1
	}
	node := this.Head.Next
	this.Head.Next = node.Next
	delete(this.Prevs, node.Key)
	if node == this.Tail {
		this.Tail = this.Head
	}
	if this.Head.Next != nil {
		this.Prevs[this.Head.Next.Key] = this.Head
	}
	return node.Key
}

func (this *LRUCache) Len() int {
	return len(this.Prevs)
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
