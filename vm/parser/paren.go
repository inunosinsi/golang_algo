package parser

import (
	"../ast"
	"../token"
)

func (p *Parser) parseParenExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}
