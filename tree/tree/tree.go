package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func New(v int) *Node {
	node := &Node{Value: v}
	return node
}

func (node *Node) Add(v int) {
	newNode := New(v)
	tree := node
	for {
		if tree.Value > newNode.Value {
			if tree.Left == nil {
				tree.Left = newNode
				break
			} else {
				tree = tree.Left
			}
		} else {
			if tree.Right == nil {
				tree.Right = newNode
				break
			} else {
				tree = tree.Right
			}
		}
	}
}

// Measure the height of a tree
func MeasureDepth(node *Node) int {
	if node.Left == nil && node.Right == nil {
		return 1
	}

	var leftHeight int
	if node.Left != nil {
		leftHeight = MeasureDepth(node.Left)
	} else {
		leftHeight = 0
	}

	var rightHeight int
	if node.Right != nil {
		rightHeight = MeasureDepth(node.Right)
	} else {
		rightHeight = 0
	}

	if leftHeight >= rightHeight {
		return 1 + leftHeight
	} else {
		return 1 + rightHeight
	}
}

func MeasureWidth(depth int) int {
	w := 1
	for i := 0; i < depth; i++ {
		w *= 2
	}
	return w
}
