package sensitive

import (
	"unicode"

	"github.com/sohaha/zlsgo/zstring"
)

type Filter struct {
	root     *trieNode
	excludes map[rune]struct{}
}

// New newFilter
func New(words ...string) *Filter {
	t := &Filter{
		root: newTrieNode(),
	}
	for _, word := range words {
		t.AddWord(word)
	}
	t.excludes = make(map[rune]struct{})
	return t
}

// AddWord append sensitive words non concurrency safe
func (t *Filter) AddWord(word string) {
	word = zstring.TrimSpace(word)
	if len(word) == 0 {
		return
	}
	node := t.root
	wChars := []rune(word)

	for _, r := range wChars {
		if unicode.IsSpace(r) {
			continue
		}
		r = ctyphe(r)

		if _, ok := node.children[r]; !ok {
			node.children[r] = newTrieNode()
		}
		node = node.children[r]
	}
	node.end = true
}

func (t *Filter) skip(r rune) bool {
	return t.inExclude(r)
}

func (t *Filter) inExclude(r rune) bool {
	_, ok := t.excludes[r]
	return ok
}

// Excludes ignore special characters between sensitive words
func (t *Filter) Excludes(items ...rune) {
	for _, item := range items {
		t.excludes[ctyphe(item)] = struct{}{}
	}
}

// Contains whether to contain sensitive words
func (t *Filter) Contains(text string) bool {
	var (
		node   *trieNode
		tChars = []rune(text)
	)

	for _, r := range tChars {
		r = ctyphe(r)

		if t.skip(r) {
			continue
		}

		if node != nil {
			node = node.getNode(r)
		}
		if node == nil {
			node = t.root.getNode(r)
		}

		if node != nil && node.end {
			return true
		}
	}
	return false
}
