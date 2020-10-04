package lexer

import (
	"fmt"

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
	s := stack.New()
	fmt.Println(s)

	return s.Stack
}

// Tape構造体に文字列をセットし、一文字目を読み始める
func new(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
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
