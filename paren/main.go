package main

import (
	"fmt"
	"log"

	"./tape"
)

var top int = 0

func main() {
	input := `if(true){
		var a = 2 * (1 + 5);
		}`
	fmt.Println("カッコの数が正しいか調べる")
	stack := tape.Read(input)

	s := new()

	for _, b := range stack {
		switch b {
		case ')':
			x := pop(s)
			if x != '(' {
				log.Fatal("()の対応が正しくありません")
			}
			break
		case '}':
			x := pop(s)
			if x != '{' {
				log.Fatal("{}の対応が正しくありません")
			}
			break
		case '(':
			push(s, b)
			break
		case '{':
			push(s, b)
			break
		default:
			//何もしない
		}
	}

	//すべて0の場合はOK
	for _, b := range s {
		if b != 0 {
			log.Fatal("error")
		}
	}

	fmt.Println("OK")
}

func new() []byte {
	return make([]byte, 10)
}

func push(s []byte, x byte) {
	s[top] = x
	top++
}

func pop(s []byte) byte {
	top--
	x := s[top]
	s[top] = 0
	return x
}
