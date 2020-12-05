package tree

import "fmt"

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

func (node *Node) Show() {
	h := MeasureHeight(node)
	fmt.Println(h)

	// asciiアートを頑張る

}

// Measure the height of a tree
func MeasureHeight(node *Node) int {
	if node.Left == nil && node.Right == nil {
		return 1
	}

	var leftHeight int
	if node.Left != nil {
		leftHeight = MeasureHeight(node.Left)
	} else {
		leftHeight = 0
	}

	var rightHeight int
	if node.Right != nil {
		rightHeight = MeasureHeight(node.Right)
	} else {
		rightHeight = 0
	}

	if leftHeight >= rightHeight {
		return 1 + leftHeight
	} else {
		return 1 + rightHeight
	}
}
