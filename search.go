package sensitive

type Position struct {
	Word  string
	Start int
	End   int
}

func (t *Filter) search(tchars []rune, handler func(start, end int)) {
	var (
		node  *trieNode
		start = -1
	)

	for i, r := range tchars {
		r = ctyphe(r)
		if t.skip(r) {
			continue
		}
		if node != nil {
			node = node.getNode(r)
		}
		if node == nil {
			start = i
			node = t.root.getNode(r)
		}
		if node != nil && node.end {
			handler(start, i+1)
			node = nil
			start = -1
		}
	}
}

// Search query all sensitive word positions
func (t *Filter) Search(text string) []Position {
	var positions []Position
	tchars := []rune(text)
	t.search(tchars, func(start, end int) {
		positions = append(positions, Position{
			Word:  string(tchars[start:end]),
			Start: start + 1,
			End:   end,
		})
	})
	return positions
}

// Replace sensitive words
func (t *Filter) Replace(text string, replace rune) string {
	tchars := []rune(text)
	t.search(tchars, func(start, end int) {
		for b := start; b < end; b++ {
			tchars[b] = replace
		}
	})
	return string(tchars)
}
