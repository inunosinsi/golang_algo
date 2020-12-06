package main

import (
	"./tree"
)

func main() {
	values := []int{9, 4, 15, 2, 6, 12, 15, 17, 3, 1, 7, 5, 8, 11, 13, 14, 16}
	bt := tree.New(values[0])
	for idx, i := range values {
		if idx == 0 {
			continue
		}
		bt.Add(i)
	}

	bt.PrintTree()
}
