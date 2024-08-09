package algorithm211

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: buildTrieNode()}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}
	node := this.root
	for _, c := range word {
		idx := c - 'a'
		if node.son[idx] == nil {
			node.son[idx] = buildTrieNode()
		}
		node = node.son[idx]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return search(word, this.root, 0)
}

func search(word string, node *TrieNode, idx int) bool {
	if idx == len(word) {
		return node.isEnd
	}
	if word[idx] == '.' {
		for i := 0; i < 26; i++ {
			if next := node.son[i]; next != nil && search(word, next, idx+1) {
				return true
			}
		}
		return false
	}
	return node.son[word[idx]-'a'] != nil && search(word, node.son[word[idx]-'a'], idx+1)
}

type TrieNode struct {
	son   []*TrieNode
	isEnd bool
}

func buildTrieNode() *TrieNode {
	return &TrieNode{son: make([]*TrieNode, 26), isEnd: false}
}
