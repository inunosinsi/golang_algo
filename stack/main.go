package main

import (
	"fmt"

	"./stack"
)

func main() {
	//スタック用のスライスを作成する
	s := stack.New()

	//スタック用スライスに値を挿入する
	s.Push(1)
	s.Push(4)
	s.Push(6)
	fmt.Println("スタックに３個の値を入れた")
	fmt.Println(s.Stack)

	//値を取り出す
	i := s.Pop()
	fmt.Println("スタックから最後に入れた値を取り出した")
	fmt.Println(i)

	fmt.Println("現在のスタックの状況")
	fmt.Println(s.Stack)

	fmt.Println("再びスタックに値を二つ入れてみる")
	s.Push(3)
	s.Push(5)

	fmt.Println("現在のスタックの状況")
	fmt.Println(s.Stack)
}
