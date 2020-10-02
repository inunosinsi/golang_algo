package stack

import "log"

var top int = 0

const capacity int = 10

//スタック用の配列を作成
func New() []int {
	return make([]int, capacity)
}

//データの挿入(push)
func Push(s []int, x int) {
	if top >= cap(s) {
		log.Fatal("スタックはオーバーフローしました")
	}
	s[top] = x
	top++
}

//データの取り出し(pop)
func Pop(s []int) int {
	if top == 0 {
		log.Fatal("スタックは空で、値の取り出しを失敗しました")
	}
	top--
	x := s[top]
	s[top] = 0
	return x
}
