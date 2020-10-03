package main

import (
	"fmt"

	"./tape"
)

func main() {
	input := "var a = 1 + 2;"
	fmt.Println("「" + input + "」から一文字ずつばらす")
	stack := tape.Read(input)
	for _, b := range stack {
		fmt.Println(string(b))
	}
}
