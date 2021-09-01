package compiler

import (
	"../ast"
)

/**
 * 1 + 2の場合
 * PUSH 1
 * PUSH 2
 * ADD
 * に変換
 */

/**
 * オペコード(Opcode)はPUSH 1のように表す
 * ニーモニック(Mnemonic)はPUSH、オペランド(Operand)は1に該当する
 */
type Opcode struct {
	Mnemonic int
	Operand  []byte
}

type Compiler struct {
	Opcodes []Opcode
}

func New() *Compiler {
	return &Compiler{}
}

//再帰でProgram → ExpressionStatement → Expressionの順で掘り下げていく
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}
	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
	}
	return nil
}
