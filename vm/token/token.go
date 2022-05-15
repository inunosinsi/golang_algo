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
	LE //<=
	GE //>=

	COMMA     //,
	SEMICOLON //;

	LPAREN   //(
	RPAREN   //)
	LBRACE   //{
	RBRACE   //}
	LBRACKET //[
	RBRACKET //]

	//キーワード
	VAR //var
	ECHO
	TRUE
	FALSE
	IF   //if
	ELSE //else
	FUNC
	RETURN

	WHILE //while

	EQ     //==
	NOT_EQ //!=
)

var keywords = map[string]int{
	"function": FUNC,
	"var":      VAR,
	"echo":     ECHO,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"while":    WHILE,
	"return":   RETURN,
}

func LookupIdent(ident string) int {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
