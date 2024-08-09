package algorithm208

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	num   int
	son   []*TrieNode
	isEnd bool
}

func buildTrieNode() *TrieNode {
	return &TrieNode{num: 1, son: make([]*TrieNode, 26), isEnd: false}
}

func Constructor() Trie {
	return Trie{root: buildTrieNode()}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	node := this.root
	for _, c := range word {
		idx := c - 'a'
		if node.son[idx] == nil {
			node.son[idx] = buildTrieNode()
		} else {
			node.son[idx].num++
		}
		node = node.son[idx]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	node := this.root
	for _, c := range word {
		if idx := c - 'a'; node.son[idx] != nil {
			node = node.son[idx]
		} else {
			return false
		}
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	node := this.root
	for _, c := range prefix {
		if idx := c - 'a'; node.son[idx] != nil {
			node = node.son[idx]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
