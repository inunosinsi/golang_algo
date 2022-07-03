package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr [5]int

	// arrの0番目のポインタ(アドレス)を格納する
	p := uintptr(unsafe.Pointer(&arr[0]))

	// arrの0番目と1番目の値のポインタを確認しておく
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])

	// 上の配列の第一引数までの距離を算出するために仮に定義した変数
	var i int

	// intのサイズ分移動したポインタを変数に代入する
	pp := (unsafe.Pointer(p + uintptr(unsafe.Sizeof(i))))
	fmt.Println(pp)

	v := (*int)(pp)

	// arrの1番目に任意の値を代入
	*v = 3
	fmt.Println(arr)
}
