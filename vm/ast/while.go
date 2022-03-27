package ast

/**
 *　while(Condition){
 *		Statements
 *	}
 */
import (
	"bytes"

	"../token"
)

type WhileExpression struct {
	Token      token.Token
	Condition  Expression
	Statements *BlockStatement
}

func (we *WhileExpression) expressionNode()      {}
func (we *WhileExpression) TokenLiteral() string { return string(we.Token.Literal) }
func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("while")
	out.WriteString(" ")
	out.WriteString(we.Condition.String())
	out.WriteString(" ")
	out.WriteString(we.Statements.String())

	return out.String()
}
