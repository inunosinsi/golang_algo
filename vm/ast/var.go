package ast

/**
 * var a = 1;のようにvarから始まる文
 */
import (
	"bytes"

	"../token"
)

type VarStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (va *VarStatement) statementNode()       {}
func (va *VarStatement) TokenLiteral() string { return string(va.Token.Literal) }

func (va *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(va.TokenLiteral() + " ")
	out.WriteString(va.Name.String())
	out.WriteString(" = ")

	if va.Value != nil {
		out.WriteString(va.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
