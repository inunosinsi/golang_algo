package main

import (
	"fmt"

	"./lexer"
	"./parser"
)

func main() {
	input := "15 + 23;"
	l := lexer.New(input)
	p := parser.New(l) //lexerをparserの中に組み込む
	program := p.Parse()
	fmt.Println(program)

	//正しい場合のコードを追加する
}
