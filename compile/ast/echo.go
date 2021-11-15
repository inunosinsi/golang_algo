package ast

import (
	"bytes"

	"../token"
)

type EchoStatement struct {
	Token token.Token
	Value Expression
}

func (ec *EchoStatement) statementNode()       {}
func (ec *EchoStatement) TokenLiteral() string { return string(ec.Token.Literal) }

func (ec *EchoStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ec.TokenLiteral() + " ")

	if ec.Value != nil {
		out.WriteString(ec.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
