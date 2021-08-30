package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 0
	}
	return n + fact(n-1)
}

func main() {
	sum := fact(7)
	fmt.Println(sum)
}
