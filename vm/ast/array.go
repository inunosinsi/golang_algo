package ast

import (
	"bytes"
	"strings"

	"../token"
)

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (*ArrayLiteral) expressionNode() {}

func (al *ArrayLiteral) TokenLiteral() string {
	if al == nil {
		return ""
	}
	return string(al.Token.Literal)
}

func (al *ArrayLiteral) String() string {
	if al == nil {
		return ""
	}

	elements := make([]string, 0, len(al.Elements))
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	var out bytes.Buffer

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (*IndexExpression) expressionNode() {}

func (ie *IndexExpression) TokenLiteral() string {
	if ie == nil {
		return ""
	}
	return string(ie.Token.Literal)
}

func (ie *IndexExpression) String() string {
	if ie == nil {
		return ""
	}

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
