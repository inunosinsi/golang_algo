package ast

/**
 *　if(Condition){
 *		Consequence(BlockStatement)
 *	}else{
 *		Alternative(BlockStatement)
 *	}
 */
import (
	"bytes"

	"../token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return string(ie.Token.Literal) }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	return out.String()
}
