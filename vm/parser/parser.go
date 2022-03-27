package parser

import (
	"../ast"
	"../lexer"
	"../token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
	blockParseFn  func() *ast.BlockStatement
)

type Parser struct {
	l *lexer.Lexer
	/**errors []string エラーメッセージは一旦保留 **/

	curToken  token.Token
	peekToken token.Token

	//map[int]...のintにTokenTypeを指定する
	prefixParseFns map[int]prefixParseFn
	infixParseFns  map[int]infixParseFn
	blockParseFns  map[int]blockParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		/**errors: []string{}, //エラーはなしにする**/
	}

	/** @register func **/
	//各parse系のメソッドを事前に登録しておく→parseExpressionで使う
	p.prefixParseFns = make(map[int]prefixParseFn)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.LPAREN, p.parseParenExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.WHILE, p.parseWhileExpression)

	p.infixParseFns = make(map[int]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LE, p.parseInfixExpression)
	p.registerInfix(token.GE, p.parseInfixExpression)

	//検証用
	p.blockParseFns = make(map[int]blockParseFn)
	p.registerBlock(token.LBRACE, p.parseBlockStatement)

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

func (p *Parser) registerPrefix(tokenType int, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType int, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerBlock(tokenType int, fn blockParseFn) {
	p.blockParseFns[tokenType] = fn
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.TokenType {
	case token.VAR:
		return p.parseVarStatement()
	case token.ECHO:
		return p.parseEchoStatement()
	case token.IDENT:
		return p.parseIdentStatement()
	case token.LBRACE: //{}の検証用
		return p.parseBlockStatement()
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
