package code

//オペコードのニーモニック
const (
	PUSH = iota
	POP  //スタックのトップの値をポップし、出力する
	ADD
	ASSIGN //スタックのトップの値をポップし、指定の変数へ書き込む
)

func GetCode(c int) string {
	switch c {
	case PUSH:
		return "PUSH"
	case POP:
		return "POP"
	case ADD:
		return "ADD"
	case ASSIGN:
		return "ASSIGN"
	}
	return ""
}
