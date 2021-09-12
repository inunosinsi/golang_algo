package parser

// prefix(前 1 + 2であれば、+の前のトークンの1)で実行する為の関数を登録する
func (p *Parser) registerPrefix(tokenType int, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// infix(前 1 + 2であれば、+のトークン)で実行する為の関数を登録する
func (p *Parser) registerInfix(tokenType int, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// 構文解析器に組み込んでいる字句解析器(lexer)のトークンを一つ進める
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// 先読み(peek)しているトークンが指定の型であるか？を調べる
func (p *Parser) peekTokenIs(tokenType int) bool {
	return p.peekToken.TokenType == tokenType
}

// 現在のトークンの優先度(precedence)を調べる
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.TokenType]; ok {
		return p
	}

	return LOWEST
}

// 先読み(peek)しているトークンの優先度(precedence)を調べる
func (p *Parser) peekPrecedence() int {
	//各トークンに設けられている優先度を取得
	if p, ok := precedences[p.peekToken.TokenType]; ok {
		return p
	}
	//指定のトークンに優先順位が設けられていなければ、最低の値を返す
	return LOWEST
}

// 先読みしているトークンが指定の型であれば、トークンの読み込みを一つ進める
func (p *Parser) expectPeek(tokenType int) bool {
	if p.peekTokenIs(tokenType) {
		p.nextToken()
		return true
	} else {
		//エラーの処理
		return false
	}
}
