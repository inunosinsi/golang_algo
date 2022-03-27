package compiler

import (
	"fmt"
	"strconv"

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
	Opcodes    []Opcode
	labelIndex int //JUMPのLn:(nには整数が入る)用の値
}

func New() *Compiler {
	return &Compiler{labelIndex: 1}
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
		case "==":
			c.emit(code.EQOP)
			return nil
		case "!=":
			c.emit(code.NEOP)
			return nil
		case "<":
			c.emit(code.LTOP)
			return nil
		case ">":
			c.emit(code.GTOP)
			return nil
		case "<=":
			c.emit(code.LEOP)
			return nil
		case ">=":
			c.emit(code.GEOP)
			return nil
		}
	case *ast.PrefixExpression:
		err := c.Compile(node.Right)
		if err != nil {
			return err
		}
		switch string(node.Operator) {
		case "!":
			c.emit(code.NOT)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}
	case *ast.IfExpression:
		err := c.Compile(node.Condition)
		if err != nil {
			return err
		}
		fjumpIndexStr := strconv.Itoa(c.labelIndex)
		c.labelIndex += 1
		c.emit(code.FJUMP, []byte("L"+fjumpIndexStr))

		err = c.Compile(node.Consequence)
		if err != nil {
			return err
		}

		if node.Alternative == nil { //elseがない場合
			c.emit(code.LABEL, []byte("L"+fjumpIndexStr+":"))
		} else { // elseがある場合
			jumpIndexStr := strconv.Itoa(c.labelIndex)
			c.labelIndex += 1
			c.emit(code.JUMP, []byte("L"+jumpIndexStr))

			c.emit(code.LABEL, []byte("L"+fjumpIndexStr+":"))
			err := c.Compile(node.Alternative)
			if err != nil {
				return err
			}
			c.emit(code.LABEL, []byte("L"+jumpIndexStr+":"))
		}
	case *ast.WhileExpression:
		jumpIndexStr := strconv.Itoa(c.labelIndex)
		c.labelIndex += 1
		c.emit(code.LABEL, []byte("L"+jumpIndexStr+":"))

		err := c.Compile(node.Condition)
		if err != nil {
			return err
		}

		fjumpIndexStr := strconv.Itoa(c.labelIndex)
		c.labelIndex += 1
		c.emit(code.FJUMP, []byte("L"+fjumpIndexStr))

		err = c.Compile(node.Statements)
		if err != nil {
			return err
		}
		c.emit(code.JUMP, []byte("L"+jumpIndexStr))
		c.emit(code.LABEL, []byte("L"+fjumpIndexStr+":"))
	case *ast.VarStatement:
		err := c.Compile(node.Value)
		if err != nil {
			return err
		}
		c.emit(code.ASSIGN, node.Name.Value)
	case *ast.BlockStatement:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}
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
	case *ast.Boolean:
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
