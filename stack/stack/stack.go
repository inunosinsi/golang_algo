package stack

import "log"

var top int = 0

const capacity int = 10

type ArrayStack struct {
	Stack []int
}

//スタック用の配列を作成
func New() *ArrayStack {
	s := &ArrayStack{Stack: make([]int, capacity)}
	return s
}

//データの挿入(push)
func (s *ArrayStack) Push(x int) {
	if top >= cap(s.Stack) {
		log.Fatal("スタックはオーバーフローしました")
	}
	s.Stack[top] = x
	top++
}

//データの取り出し(pop)
func (s *ArrayStack) Pop() int {
	if top == 0 {
		log.Fatal("スタックは空で、値の取り出しを失敗しました")
	}
	top--
	x := s.Stack[top]
	s.Stack[top] = 0
	return x
}
