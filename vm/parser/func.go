package parser

import (
	"strings"

	"../ast"
	"../token"
)

func (p *Parser) parseFunctionStatement() ast.Statement {
	fn := &ast.FunctionStatement{Token: p.curToken}

	p.nextToken()
	//関数名
	fn.Name = p.curToken.Literal

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	fn.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	fn.Body = p.parseBlockStatement()

	/** 最後にreturn文がない時は付与したい **/
	line := len(fn.Body.Statements)
	for l, stmt := range fn.Body.Statements {
		if l == line-1 {
			if strings.Index(stmt.String(), "return") < 0 {
				tk := token.Token{TokenType: token.RETURN, Literal: []byte("return")}
				stmt := &ast.ReturnStatement{
					Token: tk,
				}
				fn.Body.Statements = append(fn.Body.Statements, stmt)
			}
		}
	}

	return fn
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}
