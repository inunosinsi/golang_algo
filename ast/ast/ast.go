package ast

import "bytes"

type Node interface {
	TokenLiteral() string //TokenLiteral()があればNodeと見なす
	String() string       //動作の検証用
}

type Statement interface {
	Node
	statementNode() //statementNode()があればStatement(文)と見なす
}

type Expression interface {
	Node
	expressionNode() //expressionNode()があればExpression(式)と見なす
}

type Program struct {
	Statements []Statement //ProgramはStatement(文)が複数ある事になる
}

//動作の検証用
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
