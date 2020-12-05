package tree

import (
	"fmt"
)

var hierarchy [][]int

func (node *Node) PrintTree() {
	depth := MeasureDepth(node)

	// asciiアートを頑張る
	hierarchy = make([][]int, depth)

	for i := 0; i < depth; i++ {
		cap := MeasureWidth(i)
		hierarchy[i] = make([]int, cap)
	}
	binaryTree2Slice(node, 0, 0)

	//幅
	width := MeasureWidth(depth - 1)
	if depth < 4 {
		width += 1
	}

	for i := 0; i < len(hierarchy); i++ {
		space := drawSpace(width, i)
		fmt.Print(drawLeftSpace(width, i))
		for j := 0; j < len(hierarchy[i]); j++ {
			int := hierarchy[i][j]
			if int > 0 {
				fmt.Print(int)
			} else {
				fmt.Print("  ")
			}
			fmt.Print(space)
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

func drawLeftSpace(width int, depth int) string {
	length := calcLength(width, depth)

	s := " "

	for i := 0; i < length; i++ {
		s += " "
	}

	return s
}

func drawSpace(width int, depth int) string {
	if depth == 0 {
		return ""
	}

	length := calcLength(width, depth)

	var s string
	for i := 0; i < length; i++ {
		for j := 0; j < depth; j++ {
			s += " "
		}
	}
	s += " "

	return s
}

func calcLength(width int, depth int) int {
	length := width
	for i := 0; i < depth+1; i++ {
		if length == 1 {
			length = 0
		} else {
			length /= 2
		}
	}
	return length
}
