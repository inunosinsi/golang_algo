package parser

import (
	"../ast"
	"../token"
)

func (p *Parser) parseWhileExpression() ast.Expression {
	expression := &ast.WhileExpression{Token: p.curToken}

	// whileの後に ( があるか？
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
	expression.Statements = p.parseBlockStatement()
	return expression
}
