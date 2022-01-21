package sensitive

type trieNode struct {
	end      bool
	children map[rune]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(map[rune]*trieNode),
	}
}

func (n *trieNode) getNode(r rune) *trieNode {
	var node = n.children[r]
	return node
}
