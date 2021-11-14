package main

import (
	"fmt"

	"./hash"
	"./node"
	"./table"
)

func main() {
	t := table.New()
	n := node.New("hensu1", "cat")
	h := hash.MakeHashValue(n.Ident)
	t[h] = n

	n = node.New("hensu2", "dog")
	h = hash.MakeHashValue(n.Ident)
	t[h] = n

	result := table.Search(t, "hensu2")
	fmt.Println(result)

	// n = node.New("hensu3", "rabbit")
	// h = hash.MakeHashValue(n.Ident)
}
