package node

type Node struct {
	Ident string
	Value string
	Next  *Node
}

func New(ident string, value string) *Node {
	node := &Node{Ident: ident, Value: value}
	return node
}

func (node *Node) Add(ident string, value string) {
	newNode := New(ident, value)
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
