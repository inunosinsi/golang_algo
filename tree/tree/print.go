package tree

import "fmt"

var hierarchy [][]int

func (node *Node) PrintTree() {
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
	binaryTree2Slice(node, 0, 0)
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

func binaryTree2Slice(node *Node, depth int, n int) {
	hierarchy[depth][n] = node.Value
	if node.Left != nil {
		binaryTree2Slice(node.Left, depth+1, 2*n)
	}
	if node.Right != nil {
		binaryTree2Slice(node.Right, depth+1, 2*n+1)
	}
}

func drawSpace(i, depth int) string {
	s := " "
	for j := i; j < depth; j++ {
		for k := 0; k < depth-i-1; k++ {
			s += " "
		}
	}
	return s
}
