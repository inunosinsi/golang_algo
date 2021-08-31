package parser

import (
	"../ast"
	"../lexer"
	"../token"
)

type Parser struct {
	l *lexer.Lexer
	/**errors []string エラーメッセージは一旦保留 **/

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		/**errors: []string{}, //エラーはなしにする**/
	}

	return p
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	//処理 tokenを二回進めることで、curTokenに最初のトークン、peekTokenに２つ目のトークンが格納される
	p.nextToken()
	p.nextToken()

	//再帰下降構文解析 EOFのトークンになるまでトークンの読み込みを繰り返す
	for p.curToken.TokenType != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.TokenType {
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// 優先順位に従い、再帰を実行し続け、Expressionを完成させる
func (p *Parser) parseExpression(precedence int) ast.Expression {
	// 下の行が重要
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		switch p.curToken.TokenType {
		case token.INT:
			leftExp := p.parseIntegerLiteral()
			p.nextToken()
			//整数リテラルの後は必ず演算子がくる
			leftExp = p.parseInfixExpression(leftExp)
			return leftExp
		}
	}
	return nil
}
