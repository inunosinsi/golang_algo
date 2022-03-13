package ast

/**
 * 識別子 a = 2 + 3;であれば a のこと
 * TokenLiteral()があるので、Node型のデータになり、expressionNode()があるのでExpression型のデータでもある
 */
import (
	"bytes"

	"../token"
)

type Identifier struct {
	Token token.Token
	Value []byte
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return string(i.Token.Literal) }

//検証用
func (i *Identifier) String() string { return string(i.Value) }

/**
 * 変数に値の再代入用のノード
 */
type IdentStatement struct {
	Token      token.Token
	Value      []byte
	Expression Expression
}

func (id *IdentStatement) statementNode()       {}
func (id *IdentStatement) TokenLiteral() string { return string(id.Token.Literal) }

func (id *IdentStatement) String() string {
	var out bytes.Buffer

	out.WriteString(id.TokenLiteral() + " ")
	out.WriteString(string(id.Value))
	out.WriteString(" = ")

	if id.Expression != nil {
		out.WriteString(id.Expression.String())
	}

	out.WriteString(";")

	return out.String()
}
