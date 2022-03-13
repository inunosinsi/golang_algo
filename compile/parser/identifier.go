package parser

import (
	"../token"

	"../ast"
)

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIdentStatement() *ast.ExpressionStatement {
	//変数を定義する
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	//変数の次に = がきているか？
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
