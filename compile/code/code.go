package code

//オペコードのニーモニック
const (
	PUSH = iota
	ADD
)

func GetCode(c int) string {
	switch c {
	case PUSH:
		return "PUSH"
	case ADD:
		return "ADD"
	}
	return ""
}
