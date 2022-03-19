package parser

func (p *Parser) peekTokenIs(tokenType int) bool {
	return p.peekToken.TokenType == tokenType
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.TokenType]; ok {
		return p
	}

	return LOWEST
}

//優先順位を調べる
func (p *Parser) peekPrecedence() int {
	//各トークンに設けられている優先順位を取得
	if p, ok := precedences[p.peekToken.TokenType]; ok {
		return p
	}
	//指定のトークンに優先順位が設けられていなければ、最低の値を返す
	return LOWEST
}

//先行読み込みしているトークンが予想した(引数のt)の型であるか？
func (p *Parser) expectPeek(t int) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t int) {
	//いずれエラーメッセージを出力したい
}

func (p *Parser) curTokenIs(t int) bool {
	return p.curToken.TokenType == t
}
