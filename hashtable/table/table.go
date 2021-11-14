package table

import (
	"../hash"
	"../node"
)

func New() []*node.Node {
	return make([]*node.Node, 10, 10)
}

// 変数名で変数の値を検索する。変数テーブルに指定の変数名がなければ、文字列でnilの値を返す
func Search(t []*node.Node, ident string) string {
	//ハッシュ値を用いて検索を高速化する
	h := hash.MakeHashValue(ident)
	n := t[h]

	// 念の為に変数名が正しいか？を確認
	if n.Ident == ident {
		return n.Value
	}

	return "nil"
}
