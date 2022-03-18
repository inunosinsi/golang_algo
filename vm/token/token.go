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
	ASSIGN   //=
	PLUS     //+
	BANG     //!
	ASTERISK //*

	LT //<
	GT //>

	SEMICOLON //;

	LPAREN //(
	RPAREN //)

	//キーワード
	FNC
	VAR //var
	ECHO
	TRUE
	FALSE

	EQ     //==
	NOT_EQ //!=
)

var keywords = map[string]int{
	"function": FNC,
	"var":      VAR,
	"echo":     ECHO,
	"true":     TRUE,
	"false":    FALSE,
}

func LookupIdent(ident string) int {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
