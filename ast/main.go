package main

import (
	"fmt"

	"./lexer"
	"./parser"
)

func main() {
	input := "var a = 15 + 40;"
	l := lexer.New(input)
	p := parser.New(l) //lexerをparserの中に組み込む
	program := p.Parse()
	fmt.Println(program.String())

	//正しい場合のコードを追加する

}
