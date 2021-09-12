package parser

import (
	"strconv"

	"../ast"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	//値が正しいか？を調べる
	_, err := strconv.ParseInt(string(p.curToken.Literal), 0, 64)
	if err != nil {
		//エラーハンドリングを追加するかも
		return nil
	}

	lit.Value = p.curToken.Literal
	return lit
}
