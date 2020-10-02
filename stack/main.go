package main

import (
	"fmt"

	"./stack"
)

func main() {
	//スタック用のスライスを作成する
	s := stack.New()

	//スタック用スライスに値を挿入する
	stack.Push(s, 1)
	stack.Push(s, 4)
	stack.Push(s, 6)
	fmt.Println("スタックに３個の値を入れた")
	fmt.Println(s)

	//値を取り出す
	i := stack.Pop(s)
	fmt.Println("スタックから最後に入れた値を取り出した")
	fmt.Println(i)

	fmt.Println("現在のスタックの状況")
	fmt.Println(s)

	fmt.Println("再びスタックに値を二つ入れてみる")
	stack.Push(s, 3)
	stack.Push(s, 5)

	fmt.Println("現在のスタックの状況")
	fmt.Println(s)
}
