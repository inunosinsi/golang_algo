package main

import "fmt"

func main() {
	str := "abc"
	fmt.Println("文字列" + str + "をバイト型に変換する")

	b := []byte(str)
	fmt.Println(b)
}
