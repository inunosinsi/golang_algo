package ast

type Node interface {
	TokenLiteral() string //TokenLiteral()があればNodeと見なす
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
