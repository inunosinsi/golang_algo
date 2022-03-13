package lexer

import (
	"../stack"
	"../token"
)

type Lexer struct {
	input   string //読み込む文字列をセットする
	current int    //読み込み中の現在の位置
	next    int    //次の位置
	ch      byte   //現在読み込んでいる文字
}

//読み込んだ文字列からトークンにばらす
func Divide(input string) []token.Token {
	var tok token.Token
	s := stack.New()

	lexer := New(input)
	for {
		lexer.skip()

		switch lexer.ch {
		case '=':
			tok = newToken(token.ASSIGN, []byte("="))
		case '+':
			tok = newToken(token.PLUS, []byte("+"))
		case '*':
			tok = newToken(token.PLUS, []byte("*"))
		case ';':
			tok = newToken(token.SEMICOLON, []byte(";"))
		case '0':
			tok = newToken(token.EOF, []byte(""))
		default: //複数のbyteになりそうなもの var or 変数名 or 数字
			if isLetter(lexer.ch) {
				literal := lexer.readIdentifier()
				tokenType := token.LookupIdent(literal)
				tok = newToken(tokenType, []byte(literal))
			} else if isDigit(lexer.ch) {
				literal := lexer.readNumber()
				tok = newToken(token.INT, []byte(literal))
			} else {
				tok = newToken(token.ILLEGAL, []byte(""))
			}
		}

		s.Add(tok)

		if lexer.ch == 0 {
			break
		}

		lexer.readChar()
	}

	return s.Stack
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skip()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, []byte{l.ch})
	case '+':
		tok = newToken(token.PLUS, []byte{l.ch})
	case '*':
		tok = newToken(token.ASTERISK, []byte{l.ch})
	case ';':
		tok = newToken(token.SEMICOLON, []byte{l.ch})
	case 0:
		tok.Literal = []byte("")
		tok.TokenType = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = []byte(l.readIdentifier())
			tok.TokenType = token.LookupIdent(string(tok.Literal))
			return tok
		} else if isDigit(l.ch) {
			tok.TokenType = token.INT
			tok.Literal = []byte(l.readNumber())
			return tok
		} else {
			tok = newToken(token.ILLEGAL, []byte{l.ch})
		}
	}

	l.readChar()
	return tok
}

// Tape構造体に文字列をセットし、一文字目を読み始める
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func newToken(tokenType int, ch []byte) token.Token {
	return token.Token{TokenType: tokenType, Literal: ch}
}

// テープを一文字分進める
func (l *Lexer) readChar() {
	//文字列を最後まで読み込んだ
	if l.next >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.next]
	}
	l.current = l.next
	l.next++
}

func (l *Lexer) readIdentifier() string {
	pos := l.current
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.current]
}

func isLetter(ch byte) bool {
	return (('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_')
}

func (l *Lexer) readNumber() string {
	pos := l.current
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.current]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skip() {
	// 'a'←シングルクオートで文字列を囲うと、byte値が得られます
	for {
		if l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
			// 現在読み込んでいる文字がスペース等だった場合は一文字読み進める
			l.readChar()
		} else {
			break
		}
	}
}

func (l *Lexer) peekChar() byte {
	if l.current >= len(l.input) {
		return 0
	} else {
		return l.input[l.current]
	}
}
