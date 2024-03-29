package main

import (
	"fmt"

	"./code"
	"./compiler"
	"./lexer"
	"./parser"
)

func main() {
	//input := "echo 1;"
	input := `var a = 1 + 2;
	a = a + 1;
	echo a;`
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
