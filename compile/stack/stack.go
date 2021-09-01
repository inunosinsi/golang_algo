package stack

import "../token"

type ArrayStack struct {
	Stack []token.Token
}

const capacity = 10

var top int = 0

func New() *ArrayStack {
	s := &ArrayStack{Stack: make([]token.Token, capacity)}
	return s
}

// ArrayStackに値を加える
func (s *ArrayStack) Add(x token.Token) {
	if top >= cap(s.Stack) {
		resize(s)
	}
	s.Stack[top] = x
	top++
}

func resize(s *ArrayStack) {
	s.Stack = append(s.Stack, make([]token.Token, capacity)...)
}
