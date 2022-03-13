package parser

import (
	"../ast"
	"../token"
)

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken}

	//次のトークンがIDENT(変数に該当)であるか？を調べる
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	//変数を定義する
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	//変数の次に = がきているか？
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
