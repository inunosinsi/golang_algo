package queue

import (
	"log"
)

const capacity int = 10

var first int
var last int = capacity - 1

//キュー用の配列
func New() []int {
	return make([]int, capacity)
}

//エンキュー
func Enqueue(s []int, x int) {
	//キューが満杯になっているか？は下記の式で確かめられる
	if first > 0 && (last+1)%capacity == first {
		log.Fatal("ジョブが満杯です")
	}

	last++
	if last == capacity {
		last = 0
	}
	s[last] = x
}

//デキュー
func Dequeue(s []int) int {
	if first == last {
		log.Fatal("ジョブがありません")
	}

	x := s[first]
	s[first] = 0 //値を空にする
	first++
	if first == capacity {
		first = 0
	}
	return x
}
