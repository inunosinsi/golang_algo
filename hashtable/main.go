package main

import (
	"fmt"

	"./hash"
	"./node"
)

func main() {
	table := make([]*node.Node, 10, 10)
	n := node.New("hensu1", "cat")
	h := hash.MakeHashValue(n.Ident)
	table[h] = n

	n = node.New("hensu2", "dog")
	h = hash.MakeHashValue(n.Ident)
	table[h] = n

	result := search(table, "hensu2")
	fmt.Println(result)

	// n = node.New("hensu3", "rabbit")
	// h = hash.MakeHashValue(n.Ident)
}

// 変数名で変数の値を検索する。変数テーブルに指定の変数名がなければ、文字列でnilの値を返す
func search(table []*node.Node, ident string) string {
	//ハッシュ値を用いて検索を高速化する
	h := hash.MakeHashValue(ident)
	n := table[h]

	// 念の為に変数名が正しいか？を確認
	if n.Ident == ident {
		return n.Value
	}

	return "nil"
}
