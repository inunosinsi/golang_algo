package main

import (
	"fmt"

	"./table"
)

func main() {
	t := table.New()
	table.Add(t, "hensu1", "cat")
	table.Add(t, "hensu2", "dog")
	table.Add(t, "hensu3", "rabbit") //ハッシュ値が衝突する

	result := table.Search(t, "hensu2")
	fmt.Println(result)

	//衝突したハッシュ値の場合にどちらの検索ワードでも値にたどり着けるか？
	result = table.Search(t, "hensu3")
	fmt.Println(result)
}
