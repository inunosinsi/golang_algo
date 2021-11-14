package main

import (
	"fmt"

	"./node"
	"./table"
)

func main() {
	// 変数名と変数に代入した文字列を格納するテーブルをスライス(可変長の配列)で用意する
	t := table.New()
	t[0] = node.New("hensu1", "cat")
	t[1] = node.New("hensu2", "dog")
	t[2] = node.New("hensu3", "rabbit")
	t[3] = node.New("hensu4", "turtle")
	t[4] = node.New("hensu5", "bird")
	t[5] = node.New("hensu6", "crab")

	// 変数hensu4に格納された値を調べる
	result := table.Search(t, "hensu4")
	fmt.Println(result)
}
