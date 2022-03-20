package ast

/**
 * !trueや-1の!や-の何らかの値の前に付与する符号
 */
import (
	"bytes"

	"../token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator []byte
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return string(pe.Token.Literal) }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(string(pe.Operator))
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
