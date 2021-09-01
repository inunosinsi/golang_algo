package ast

/**
 * 識別子 a = 2 + 3;であれば a のこと
 * TokenLiteral()があるので、Node型のデータになり、expressionNode()があるのでExpression型のデータでもある
 */
import "../token"

type Identifier struct {
	Token token.Token
	Value []byte
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return string(i.Token.Literal) }

//検証用
func (i *Identifier) String() string { return string(i.Value) }
