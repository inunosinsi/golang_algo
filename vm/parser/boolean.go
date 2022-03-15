package parser

import (
	"../ast"
)

func (p *Parser) parseBoolean() ast.Expression {
	var v []byte
	if string(p.curToken.Literal) == "true" {
		v = []byte("1")
	} else {
		v = []byte("0")
	}
	return &ast.Boolean{Token: p.curToken, Value: v}
}
