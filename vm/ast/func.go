package ast

import (
	"bytes"
	"strings"

	"../token"
)

type FunctionStatement struct {
	Token      token.Token
	Name       []byte
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionStatement) statementNode()       {}
func (fl *FunctionStatement) TokenLiteral() string { return string(fl.Token.Literal) }
func (fl *FunctionStatement) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(string(fl.Name))
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString("{ ")
	out.WriteString(fl.Body.String())
	out.WriteString(" }")

	return out.String()
}
