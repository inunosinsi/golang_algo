package ast

/**
 * 2 + 3の場合、下記の木構造にする
 *     +
 *    / \
 *   /   \
 *  2     3
 */
import "../token"

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator []byte
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return string(oe.Token.Literal) }
