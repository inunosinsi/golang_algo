package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	var slice []int
	slice = make([]int, 10, 20)
	slice[1] = 5
	for i, _ := range slice {
		fmt.Println(&slice[i])
	}
}
