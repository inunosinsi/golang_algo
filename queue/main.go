package main

import (
	"fmt"

	"./queue"
)

func main() {
	s := queue.New()

	//オーバーフローの準備 値を10個
	for i := 1; i <= 10; i++ {
		queue.Enqueue(s, i)
	}

	fmt.Println("値がすべて埋まっている状態")
	fmt.Println(s)

	fmt.Println("最初にいれた1を取り出す")
	fmt.Println(queue.Dequeue(s))

	fmt.Println("リングバッファになっているか確認するため、新たに値を入れてみる")
	queue.Enqueue(s, 11)

	fmt.Println("値がすべて埋まっている状態")
	fmt.Println(s)

	fmt.Println("値を追加してオーバーフローにする")
	queue.Enqueue(s, 12)

	fmt.Println(s)
}
