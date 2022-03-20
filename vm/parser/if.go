package parser

import (
	"../ast"
	"../token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	// ifの後に ( があるか？
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	// if文の条件式が終わった後に ) があるか？
	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	// ) の後に { はあるか？
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// { の後はBlockStatementになる
	expression.Consequence = p.parseBlockStatement()
	return expression
}
