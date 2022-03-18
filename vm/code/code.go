package code

//オペコードのニーモニック
const (
	PUSH = iota
	POP  //スタックのトップの値をポップし、出力する
	ADD
	NOT //!で真偽反転
	MUL
	ASSIGN //スタックのトップの値をポップし、指定の変数へ書き込む
	EQOP   //==
	NEOP   //!=
	LTOP   //<
	GTOP   //>
	LEOP   //<=
	GEOP   //>=
)

func GetCode(c int) string {
	switch c {
	case PUSH:
		return "PUSH"
	case POP:
		return "POP"
	case ADD:
		return "ADD"
	case NOT:
		return "NOT"
	case MUL:
		return "MUL"
	case ASSIGN:
		return "ASSIGN"
	case EQOP:
		return "EQOP"
	case NEOP:
		return "NEOP"
	case LTOP:
		return "LTOP"
	case GTOP:
		return "GTOP"
	case LEOP:
		return "LEOP"
	case GEOP:
		return "GEOP"
	}
	return ""
}
