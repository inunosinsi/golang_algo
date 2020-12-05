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

var hierarchy [][]int

func (node *Node) Show() {
	depth := MeasureDepth(node)

	// asciiアートを頑張る
	hierarchy = make([][]int, depth)

	for i := 0; i < depth; i++ {
		cap := 1
		for j := 0; j < i; j++ {
			cap *= 2
		}
		hierarchy[i] = make([]int, cap)
	}
	bt2arr(node, 0, 0)
	//fmt.Println(hierarchy)
	for i := 0; i < len(hierarchy); i++ {
		for j := 0; j < len(hierarchy[i]); j++ {
			int := hierarchy[i][j]
			fmt.Print(drawSpace(i, depth))
			if int > 0 {
				fmt.Print(int)
			} else {
				fmt.Print("  ")
			}
			fmt.Print(drawSpace(i, depth))
		}
		fmt.Print("\n")
	}
}

func bt2arr(node *Node, depth int, n int) {
	hierarchy[depth][n] = node.Value
	if node.Left != nil {
		bt2arr(node.Left, depth+1, 2*n)
	}
	if node.Right != nil {
		bt2arr(node.Right, depth+1, 2*n+1)
	}
}

func drawSpace(i, depth int) string {
	space := " "
	for j := i; j < depth; j++ {
		for k := 0; k < depth-i-1; k++ {
			space += " "
		}
	}
	return space
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
