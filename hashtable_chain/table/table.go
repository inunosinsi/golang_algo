package table

import (
	"../hash"
	"../node"
)

func New() []*node.Node {
	return make([]*node.Node, 10, 10)
}

func Add(t []*node.Node, ident string, value string) []*node.Node {
	h := hash.MakeHashValue(ident)
	newNode := node.New(ident, value)
	n := t[h]

	// nがnilの場合はtableにそのまま加える
	if n == nil {
		t[h] = newNode
		return t
	}

	//ハッシュ値の箇所に既に値がある場合はNodeをNextに数珠つなぎする
	if len(n.Ident) > 0 {
		for {
			if n.Next == nil {
				n.Next = newNode
				break
			}
			n = n.Next
		}
	}
	return t
}

// 変数名で変数の値を検索する。変数テーブルに指定の変数名がなければ、文字列でnilの値を返す
func Search(t []*node.Node, ident string) string {
	//ハッシュ値を用いて検索を高速化する
	h := hash.MakeHashValue(ident)
	n := t[h]

	// 念の為に変数名が正しいか？を確認
	for {
		if n.Ident == ident {
			return n.Value
		}
		if n.Next == nil {
			break
		}
		n = n.Next
	}

	return "nil"
}
