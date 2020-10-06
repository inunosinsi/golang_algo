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

	lexer := new(input)
	for {
		lexer.skip()

		switch lexer.ch {
		case '=':
			tok = newToken(token.ASSIGN, []byte("="))
			break
		case '0':
			tok = newToken(token.EOF, []byte(""))
		default: //複数のbyteになりそうなもの var or 変数名 or 数字
			tok = newToken(token.ILLEGAL, []byte(""))
		}

		//fmt.Println(lexer.ch)
		s.Add(tok)

		if lexer.ch == 0 {
			break
		}

		lexer.readChar()
	}

	return s.Stack
}

// Tape構造体に文字列をセットし、一文字目を読み始める
func new(input string) *Lexer {
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
	if l.current >= len(l.input) {
		l.ch = 0
		l.current = 0
		return
	}

	l.ch = l.input[l.current]
	l.current++
}

func (l *Lexer) skip() {
	// 'a'←シングルクオートで文字列を囲うと、byte値が得られます
	for {
		if l.ch == ' ' || l.ch == '\t' {
			// 現在読み込んでいる文字がスペース等だった場合は一文字読み進める
			l.readChar()
		} else {
			break
		}
	}
}
