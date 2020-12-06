package tree

/****************************************************************************************************************************************************************/
/* PrintTreeは下記のページを参考にして作成                                                                                                                           */
/* https://ja.stackoverflow.com/questions/4739/2-%E5%88%86%E6%9C%A8%E3%82%92-ascii-art-%E3%81%A7%E8%A1%A8%E7%A4%BA%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95%E3%81%AF */
/****************************************************************************************************************************************************************/

import "strconv"

type TreeToStringFunc func(*Node) string

// Prints the given binary tree
func (node *Node) PrintTree() {
	printNodeWithPrinter(node, defaultTreeToStringFunc())
}

func defaultTreeToStringFunc() TreeToStringFunc {
	return func(node *Node) string { return strconv.Itoa(node.Value) }
}

func printNodeWithPrinter(node *Node, treeToString TreeToStringFunc) {
	printNodeInternal([]*Node{node}, 1, maxLevel(node), digits(maxValue(node))/2-1, treeToString)
}

func printNodeInternal(nodes []*Node, level int, maxLevel int, floorElev int, treeToString TreeToStringFunc) {
	if len(nodes) == 0 || isElementsNull(nodes) {
		return
	}

	floor := maxLevel - level + floorElev + 1
	edgeLines := pow(2, max(floor-1, 0))
	firstSpaces := pow(2, floor) - 1
	betweenSpaces := pow(2, floor+1) - 1

	var newNodes []*Node
	var prevPostDigits = 0
	for i, node := range nodes {
		var relativeLeft, relativeRight *Node
		var value string
		var digits int
		if node != nil {
			relativeLeft, relativeRight = node.Left, node.Right
			value = treeToString(node)
		} else {
			relativeLeft, relativeRight = nil, nil
			value = " "
		}
		digits = len(value)

		var preDigits, postDigits = (digits / 2), (digits / 2) + (digits % 2)

		if i == 0 {
			printWhitespaces(firstSpaces - (preDigits + prevPostDigits))
		} else {
			printWhitespaces(betweenSpaces - (preDigits + prevPostDigits) + 1)
		}
		newNodes = append(newNodes, relativeLeft)
		newNodes = append(newNodes, relativeRight)
		print(value)

		prevPostDigits = postDigits
	}
	println("")

	startEdgeLine := edgeLines
	if maxLevel != level-1 {
		startEdgeLine = startEdgeLine - 1
	}
	for i := startEdgeLine; i <= edgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			printWhitespaces(firstSpaces - i)
			if nodes[j] == nil {
				printWhitespaces(edgeLines + edgeLines + i + 1)
				continue
			}
			if i == startEdgeLine {
				printTreeBranchUpper(nodes[j], i)
			} else {
				printTreeBranchLower(nodes[j], i)
			}
			printWhitespaces(edgeLines + edgeLines - i)
		}
		println("")
	}
	printNodeInternal(newNodes, level+1, maxLevel, floorElev, treeToString)
}

func getPrintTreeBranchFunc(printLeft func(int), printCenter func(int), printRight func(int),
	printSpace func(int), leftistLetter rune, rightestLetter rune) func(*Node, int) {
	return func(node *Node, i int) {
		if node.Left == nil {
			printSpace(i)
		} else if i > 0 {
			print(string(leftistLetter))
			printLeft(i - 1)
		}
		if node.Left == nil && node.Right == nil {
			printSpace(1)
		} else {
			printCenter(1)
		}
		if node.Right == nil {
			printSpace(i)
		} else if i > 0 {
			printRight(i - 1)
			print(string(rightestLetter))
		}
	}
}

func getPrintLetters(letter rune) func(int) {
	return func(count int) {
		for i := 0; i < count; i++ {
			print(string(letter))
		}
	}
}

var printWhitespaces = getPrintLetters(rune(' '))
var printHorizontalLine = getPrintLetters(rune('-'))
var printVirticalLine = getPrintLetters(rune('|'))
var printTreeBranchUpper = getPrintTreeBranchFunc(
	printHorizontalLine, printVirticalLine, printHorizontalLine, printWhitespaces,
	rune('-'), rune('-'))
var printTreeBranchLower = getPrintTreeBranchFunc(
	printWhitespaces, printWhitespaces, printWhitespaces, printWhitespaces,
	rune('/'), rune('\\'))

func maxLevel(node *Node) int {
	if node == nil {
		return 0
	}
	return max(maxLevel(node.Left), maxLevel(node.Right)) + 1
}

func maxValue(node *Node) int {
	mx := node.Value
	if node.Left != nil {
		mx = max(mx, maxValue(node.Left))
	}
	if node.Right != nil {
		mx = max(mx, maxValue(node.Right))
	}
	return mx
}

func isElementsNull(list []*Node) bool {
	for _, object := range list {
		if object != nil {
			return false
		}
	}
	return true
}

//
// Math Utilities for Integer (Pow, Max for int)
//

func pow(x, y int) (r int) {
	if x == r || y < r {
		return
	}
	r = 1
	if x == r {
		return
	}
	if x < 0 {
		x = -x
		if y&r == r {
			r = -r
		}
	}
	for y > 0 {
		if y&1 == 1 {
			r *= x
		}
		x *= x
		y >>= 1
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func digits(a int) int {
	if a == 0 {
		return 1
	}
	var c int
	for c = 0; a != 0; c++ {
		a /= 10
	}
	return c
}
