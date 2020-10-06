package token

//トークン
type Token struct {
	TokenType int
	Literal   []byte
}

const (
	//
	EOF     = iota
	ILLEGAL //わからないものはすべてこれ

	//演算子
	ASSIGN //=
	PLUS   //+

	//キーワード
	VAR //var
)
