package ast

/**
 * 2 + 3の場合、下記の木構造にする
 *     +
 *    / \
 *   /   \
 *  2     3
 */
import (
	"bytes"

	"../token"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator []byte
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return string(oe.Token.Literal) }

func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + string(oe.Operator) + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}
