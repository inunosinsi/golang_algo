package main

import (
	"fmt"

	"./code"
	"./compiler"
	"./lexer"
	"./parser"
)

func main() {
	input := "var a = 15 + 40;"
	l := lexer.New(input)
	p := parser.New(l) //lexerをparserの中に組み込む
	program := p.Parse()
	c := compiler.New()
	_ = c.Compile(program) //エラーを拾う事を無しにする

	// 中間コードを出力する
	if len(c.Opcodes) > 0 {
		for _, opcode := range c.Opcodes {
			fmt.Printf("%s %s\n", code.GetCode(opcode.Mnemonic), string(opcode.Operand))
		}
	}
}
