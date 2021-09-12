package parser

import (
	"../ast"
	"../lexer"
	"../token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer
	/**errors []string エラーメッセージは一旦保留 **/

	// 構文解析器に組み込んだ字句解析器が現在読み込んでいるトークン
	curToken token.Token
	// 構文解析器に組み込んだ字句解析器が先読みしているトークン
	peekToken token.Token

	//map[int]...のintにTokenTypeを指定する
	prefixParseFns map[int]prefixParseFn
	infixParseFns  map[int]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		/**errors: []string{}, //エラーはなしにする**/
	}

	/** @register func **/
	//各parse系のメソッドを事前に登録しておく→parseExpressionで使う
	p.prefixParseFns = make(map[int]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	p.infixParseFns = make(map[int]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)

	//処理 tokenを二回進めることで、curTokenに最初のトークン、peekTokenに２つ目のトークンが格納される
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

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

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.TokenType {
	case token.VAR:
		return p.parseVarStatement()
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
	//@register funcで登録したメソッドを呼び出す
	prefix := p.prefixParseFns[p.curToken.TokenType]
	leftExp := prefix()

	//再帰下降構文解析の要　@register funcで登録したメソッドを呼び出す
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.TokenType]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}
