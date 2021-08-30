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
		/**errors: []string{},**/
	}

	return p
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	//処理

	return program
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// func (p *Parser) parseProgram() *ast.Program {
//
// }
//
// func (p *Parser) parseExpression() ast.Expression {
//
// }
