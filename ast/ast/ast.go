package ast

//TokenLiteral()があればNodeと見なす
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode() //
}

type Program struct {
	Statements []Statement
}
