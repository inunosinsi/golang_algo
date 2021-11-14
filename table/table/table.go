package table

import (
	"../node"
)

func New() []*node.Node {
	return make([]*node.Node, 10, 10)
}

// 変数名で変数の値を検索する。変数テーブルに指定の変数名がなければ、文字列でnilの値を返す
func Search(t []*node.Node, ident string) string {
	for _, node := range t {
		if node.Ident == ident {
			return node.Value
		}
	}
	return "nil"
}
