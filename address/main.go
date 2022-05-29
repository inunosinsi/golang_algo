package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := []int{1, 2}
	p := uintptr(unsafe.Pointer(&arr[0])) // arrの0番目のポインタ(アドレス)を格納する
	fmt.Println(&arr[0])
	fmt.Println(&arr[1]) // arrの1番目の値のポインタを確認しておく

	var i int                                             // 上の配列の第一引数までの距離を算出するために仮に定義した変数
	pp := (unsafe.Pointer(p + uintptr(unsafe.Sizeof(i)))) // intのサイズ分移動したポインタを変数に代入する
	fmt.Println(pp)

	v := (*int)(pp)

	// arrの1番目の値を出力する
	fmt.Println(*v)

	// arrの1番目に任意の値を代入
	*v = 3
	fmt.Println(arr)
}
