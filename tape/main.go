package main

import (
	"fmt"

	"./tape"
)

func main() {
	list := tape.Read("var a = 1 + 2;")
	for _, b := range list {
		fmt.Println(string(b))
	}
}
