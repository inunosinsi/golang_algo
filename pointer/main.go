package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)

	//1番目の要素のアドレスを取得
	p := &arr[1]
	fmt.Println(p)

	*p = 5
	fmt.Println(arr)
}
