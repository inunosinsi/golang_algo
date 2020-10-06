package main

import (
	"fmt"

	"./lexer"
)

func main() {
	input := "var a = 1 + 2;"
	fmt.Println("「" + input + "」からトークンにばらす")
	stack := lexer.Divide(input)
	for _, token := range stack {
		fmt.Println(string(token.Literal))
	}
}
