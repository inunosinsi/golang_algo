package tape

import (
	"../stack"
)

// 指定の文字列を1文字ずつ読む為のテープ
type Tape struct {
	input   string //読み込む文字列をセットする
	current int    //読み込み中の現在の位置
	ch      byte   //現在読み込んでいる文字
}

// 位置文字ずつ取り出し、byteの配列に格納する
func Read(input string) []byte {
	s := stack.New()

	tape := new(input)
	for {
		//最初にスペース等を除く @ToDo 課題にする
		tape.skip()

		s.Add(tape.ch)

		// テープを最後まで読み進めるとchは0になるので、0でテープの読み込みを終了する
		if tape.ch == 0 {
			break
		}
		tape.readChar()
	}

	return s.Stack
}

// Tape構造体に文字列をセットし、一文字目を読み始める
func new(input string) *Tape {
	t := &Tape{input: input}
	t.readChar()
	return t
}

// テープを一文字分進める
func (t *Tape) readChar() {
	//文字列を最後まで読み込んだ
	if t.current >= len(t.input) {
		t.ch = 0
		t.current = 0
		return
	}

	t.ch = t.input[t.current]
	t.current++
}

// スペースを除く　@ToDo 課題 どこに入れるを機能するか考えてみる
func (t *Tape) skip() {
	// 'a'←シングルクオートで文字列を囲うと、byte値が得られます
	for {
		if t.ch == ' ' || t.ch == '\t' || t.ch == '\r' || t.ch == '\n' {
			// 現在読み込んでいる文字がスペース等だった場合は一文字読み進める
			t.readChar()
		} else {
			break
		}
	}
}
