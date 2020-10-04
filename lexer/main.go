package main

import (
	"fmt"

	"./lexer"
)

func main() {
	input := "var a = 1 + 2;"
	fmt.Println("「" + input + "」から一文字ずつばらす")
	stack := lexer.Divide(input)
	for _, token := range stack {
		fmt.Println(string(token.Literal))
	}
}
