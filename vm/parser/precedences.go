package parser

import "../token"

/** 優先順位 再帰下降構文解析の要 **/
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > または <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X または!X
	CALL        // myFunction(X)
	INDEX       // array[index]
)

// map[int]intの[int]はtoken.TokenTypeに対応
var precedences = map[int]int{
	token.EQ:     EQUALS,
	token.NOT_EQ: EQUALS,
	token.LT:     LESSGREATER,
	token.GT:     LESSGREATER,
	token.PLUS:   SUM,
	// token.MINUS:    SUM,
	// token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.LPAREN:   CALL,
	// token.LBRACKET: INDEX,
}

/** 優先順位 **/
