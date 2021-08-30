package main

import (
	"fmt"

	"./lexer"
)

func main() {
	input := "15 + 23;"
	stack := lexer.Divide(input)
	// for _, token := range stack {
	// 	fmt.Println(string(token.Literal))
	// }

	fmt.Println(stack)
	//stackからastを作る
}
