package main

import (
	"./tree"
)

func main() {
	t := tree.New(9)
	t.Add(4)
	t.Add(15)
	t.Add(2)
	t.Add(6)
	t.Add(12)
	t.Add(17)

	t.Show()
}
