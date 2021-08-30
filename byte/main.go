package main

import (
	"fmt"
	"strconv"
)

func main() {
	var b byte
	b = '9' // int型の1をシングルクオートで囲み、byte型に変える

	i := byte2int(b)
	fmt.Println(i)
}

// byte型からint型に変換する
func byte2int(b byte) int {
	//念の為にbyteの値が1桁の整数であるか調べる
	if b < '0' || b > '9' {
		return 0
	}

	//string関数を利用する為、byte型の値を[]byte型に変換する
	bytes := []byte{b}
	str := string(bytes)

	//strconv.Atoiでstring型からint型に変換する
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}
