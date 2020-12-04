package sllist

type Node struct {
	Value int
	Next  *Node
}

func New(v int) *Node {
	node := &Node{Value: v}
	return node
}

func (node *Node) Add(v int) {
	newNode := New(v)
	list := node
	//Nextが空のNodeまで辿る
	for {
		if list.Next == nil {
			list.Next = newNode
			break
		}
		list = list.Next
	}
}
