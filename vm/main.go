package main

import (
	"fmt"

	"./code"
	"./compiler"
	"./lexer"
	"./parser"
)

func main() {
	input := `var a = 0;
while(a < 10){
	a = a + 1;
	echo a;
}`
	l := lexer.New(input)
	p := parser.New(l) //lexerをparserの中に組み込む
	program := p.Parse()

	//抽象構文木が正しくできれば出力される
	fmt.Println(program.String())

	c := compiler.New()
	_ = c.Compile(program) //エラーを拾う事を無しにする

	// 中間コードを出力する
	if len(c.Opcodes) > 0 {
		for _, opcode := range c.Opcodes {
			if opcode.Mnemonic == code.LABEL {
				fmt.Printf("%s\n", string(opcode.Operand))
			} else {
				fmt.Printf("%s %s\n", code.GetCode(opcode.Mnemonic), string(opcode.Operand))
			}
		}
	}

	// result := vm.Eval(c.Opcodes)
	// fmt.Println(result)
}
