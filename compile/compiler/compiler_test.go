package compiler

import (
	"strings"
	"testing"

	"../code"
	"../lexer"
	"../parser"
)

type compilerTestCase struct {
	input                    string //入力するコード
	expectedIntermediateCode string //生成される正しい中間コード
}

func TestCompiler(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: "1 + 2",
			expectedIntermediateCode: `PUSH 1
PUSH 2
ADD`,
		},
		{
			input: "var a = 1 + 2",
			expectedIntermediateCode: `PUSH 1
PUSH 2
ADD
ASSIGN a`,
		},
		{
			input: `var a = 1 + 2;
echo a;`,
			expectedIntermediateCode: `PUSH 1
PUSH 2
ADD
ASSIGN a
PUSH a
POP`,
		},
	}

	for _, tt := range tests {
		lines := generate(tt.input)
		exp := trimString(tt.expectedIntermediateCode)

		if lines != exp {
			t.Errorf("compiler error : \"" + tt.input + "\"")
		}
	}

}

func generate(ipt string) string {
	l := lexer.New(ipt)
	p := parser.New(l) //lexerをparserの中に組み込む
	program := p.Parse()
	c := New()
	_ = c.Compile(program) //エラーを拾う事を無しにする

	lines := ""
	if len(c.Opcodes) > 0 {
		for _, opcode := range c.Opcodes {
			m := code.GetCode(opcode.Mnemonic)
			o := string(opcode.Operand)
			lines += trimString(m+" "+o) + "\n"
		}
	}
	return trimString(lines)
}

func trimString(lines string) string {
	lines = strings.Trim(lines, "\n")
	lines = strings.Trim(lines, " ")
	return lines
}
