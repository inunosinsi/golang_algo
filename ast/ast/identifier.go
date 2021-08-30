package ast

/**
 * 識別子 a = 2 + 3;であれば a のこと
 * TokenLiteral()があるので、Node型のデータになり、expressionNode()があるのでExpression型のデータでもある
 */
import "../token"

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return string(i.Token.Literal) }
