package main

import "fmt"

func f(n int) int {
	if n < 2 {
		return n
	}
	return f(n-2) + f(n-1)
}

func main() {
	r := f(8)
	fmt.Println(r)
}
