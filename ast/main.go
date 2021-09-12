package main

import (
	"fmt"

	"./lexer"
	"./parser"
)

func main() {
	input := "15 + 40;"
	l := lexer.New(input)
	p := parser.New(l) //lexerをparserの中に組み込む
	program := p.Parse()

	//抽象構文木が正しくできれば出力される
	fmt.Println(program.String())
}
