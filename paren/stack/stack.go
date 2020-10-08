package stack

type ArrayStack struct {
	Stack []byte
}

const capacity = 10

var top int = 0

func New() *ArrayStack {
	s := &ArrayStack{Stack: make([]byte, capacity)}
	return s
}

// ArrayStackに値を加える
func (s *ArrayStack) Add(x byte) {
	if top >= len(s.Stack) {
		resize(s)
	}
	s.Stack[top] = x
	top++
}

func resize(s *ArrayStack) {
	s.Stack = append(s.Stack, make([]byte, capacity)...)
}
