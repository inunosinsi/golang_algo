package token

//トークン
type Token struct {
	TokenType int
	Literal   []byte
}

const (
	//
	EOF = iota
)
