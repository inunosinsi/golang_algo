package ast

/**
 * 数字 a = 2 + 3;であれば、2と3のこと
 */
import "../token"

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return string(il.Token.Literal) }
