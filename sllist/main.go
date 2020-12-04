package main

import (
	"fmt"

	"./sllist"
)

func main() {
	list := sllist.New(2)

	//値を3個加えてみる
	list.Add(5)
	list.Add(3)
	list.Add(8)

	//値が4個あるか？確認してみる
	node := list
	for {
		fmt.Println(node.Value)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}
