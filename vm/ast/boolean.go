package ast

/**
 * booleanは真(true)と偽(false)のこと
 * trueは1, falseは0とする
 */
import (
	"../token"
)

type Boolean struct {
	Token token.Token
	Value []byte
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return string(b.Token.Literal) }

//検証用
func (b *Boolean) String() string {
	if string(b.Value) == "1" {
		return "true"
	} else {
		return "false"
	}
}
