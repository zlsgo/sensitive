package sensitive

import (
	"github.com/sohaha/zlsgo/zutil"
)

func (t *Filter) find(tchars []rune, handler func(str string) bool) {
	var (
		node *trieNode
		b    = zutil.GetBuff()
	)

	defer zutil.PutBuff(b)

	for _, r := range tchars {
		nr := ctyphe(r)
		if t.skip(nr) {
			if node != nil {
				b.WriteRune(r)
			} else {
				b.Reset()
			}
			continue
		}
		if node != nil {
			node = node.getNode(nr)
		}
		if node == nil {
			b.Reset()
			node = t.root.getNode(nr)
		}
		b.WriteRune(r)
		if node != nil && node.end {
			if !handler(b.String()) {
				return
			}
			node = nil
			b.Reset()
		}
	}
}

func (t *Filter) FindAll(text string) []string {
	tchars := []rune(text)
	matchText := make([]string, 0)
	t.find(tchars, func(str string) bool {
		matchText = append(matchText, str)
		return true
	})
	return matchText
}

func (t *Filter) First(text string) string {
	tchars := []rune(text)
	matchText := ""
	t.find(tchars, func(str string) bool {
		matchText = str
		return false
	})

	return matchText
}
