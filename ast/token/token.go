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

	//識別子 + リテラル
	IDENT
	INT

	//演算子
	ASSIGN //=
	PLUS   //+

	ASTERISK  //*
	SEMICOLON //;

	//キーワード
	FNC
	VAR //var
)

var keywords = map[string]int{
	"function": FNC,
	"var":      VAR,
}

func LookupIdent(ident string) int {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
