package main

import (
	"fmt"

	"./compiler"
	"./lexer"
	"./parser"
)

func main() {
	input := "15 + 40 + 50;"
	l := lexer.New(input)
	p := parser.New(l) //lexerをparserの中に組み込む
	program := p.Parse()
	c := compiler.New()
	opcodes := c.Compile(program)
	fmt.Println(opcodes)

	//正しい場合のコードを追加する

}
