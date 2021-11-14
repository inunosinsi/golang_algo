package main

import (
	"fmt"

	"./node"
)

func main() {
	// 変数名と変数に代入した文字列を格納するテーブルをスライス(可変長の配列)で用意する
	table := make([]*node.Node, 10, 10)
	table[0] = node.New("hensu1", "cat")
	table[1] = node.New("hensu2", "dog")
	table[2] = node.New("hensu3", "rabbit")
	table[3] = node.New("hensu4", "turtle")
	table[4] = node.New("hensu5", "bird")
	table[5] = node.New("hensu6", "crab")

	// 変数hensu4に格納された値を調べる
	result := search(table, "hensu4")
	fmt.Println(result)
}

// 変数名で変数の値を検索する。変数テーブルに指定の変数名がなければ、文字列でnilの値を返す
func search(table []*node.Node, ident string) string {
	for _, node := range table {
		if node.Ident == ident {
			return node.Value
		}
	}
	return "nil"
}
