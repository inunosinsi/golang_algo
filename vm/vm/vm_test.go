package vm

import (
	"testing"

	"../compiler"
	"../lexer"
	"../parser"
)

type vmTestCase struct {
	input                     string //入力するコード
	expectedIntermediateValue int    //結果
}

func TestEval(t *testing.T) {
	tests := []vmTestCase{
		{
			input:                     "echo 1;",
			expectedIntermediateValue: 1,
		},
		{
			input: `var a = 2 + 1;
			a = a + 1;
			echo a;`,
			expectedIntermediateValue: 4,
		},
	}

	for _, tt := range tests {
		codes := generate(tt.input)
		if Eval(codes) != tt.expectedIntermediateValue {
			t.Errorf("vm error : \"" + tt.input + "\"")
		}
	}
}

func generate(ipt string) []compiler.Opcode {
	l := lexer.New(ipt)
	p := parser.New(l) //lexerをparserの中に組み込む
	program := p.Parse()
	c := compiler.New()
	_ = c.Compile(program) //エラーを拾う事を無しにする
	return c.Opcodes
}
