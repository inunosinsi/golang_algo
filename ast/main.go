package main

import (
	"./lexer"
	"./parser"
)

func main() {
	input := "15 + 23;"
	l := lexer.New(input)
	p := parser.New(l) //lexerをparserの中に組み込む
	p.Parse()
}
