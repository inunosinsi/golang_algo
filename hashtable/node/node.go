package node

type Node struct {
	Ident string
	Value string
}

func New(ident string, value string) *Node {
	node := &Node{Ident: ident, Value: value}
	return node
}
