package compiler

import (
	"../ast"
	"../code"
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
	case *ast.IntegerLiteral:
		c.emit(code.PUSH, node.Value)
	case *ast.InfixExpression:
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}

		err = c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch string(node.Operator) {
		case "+":
			c.emit(code.ADD)
			return nil
		case "*":
			c.emit(code.MUL)
			return nil
		}
	case *ast.VarStatement:
		err := c.Compile(node.Value)
		if err != nil {
			return err
		}
		c.emit(code.ASSIGN, node.Name.Value)
	case *ast.IdentStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
		c.emit(code.ASSIGN, node.Value)
	case *ast.EchoStatement:
		err := c.Compile(node.Value)
		if err != nil {
			return err
		}
		c.emit(code.POP)
	case *ast.Identifier:
		c.emit(code.PUSH, node.Value)
	}

	return nil
}

func (c *Compiler) emit(mnemonic int, operands ...[]byte) {
	op := Opcode{
		Mnemonic: mnemonic,
	}

	if len(operands) > 0 {
		op.Operand = operands[0]
	}

	c.Opcodes = append(c.Opcodes, op)
}
