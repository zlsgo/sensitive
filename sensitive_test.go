package sensitive_test

import (
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/sensitive"
)

var (
	example = "他是个傻.B啦，你好吗？滚啦！白痴! 笨-蛋! Ｆooｌ! ．．．"
)

func TestBase(t *testing.T) {
	tt := zlsgo.NewTest(t)
	s := sensitive.New("傻B啦", "滚", "笨蛋", "fool")

	s.Excludes('-', '.')

	first := s.First(example)
	t.Log(first)

	findWords := s.FindAll(example)
	t.Log(findWords)

	exist := s.Contains(example)
	tt.EqualTrue(exist)
	exist = s.Contains("example")
	tt.EqualTrue(!exist)

	positions := s.Search(example)
	t.Log(example, positions)

	newText := s.Replace(example, '*')
	t.Log(newText)
}

func TestAdd(t *testing.T) {
	s := sensitive.New()

	s.Excludes('-')

	s.AddWord(" ")
	s.AddWord("白痴")
	s.AddWord("笨蛋")
	s.AddWord("fool")

	findWords := s.FindAll(example)
	t.Log(findWords)
}
