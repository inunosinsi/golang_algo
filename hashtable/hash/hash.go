package hash

import (
	"crypto/md5"
	"strconv"
)

//変数名からハッシュ値を作成する。コードの都合上、ハッシュ値は0〜9までの整数型にする
func MakeHashValue(ident string) int {
	hash := md5.Sum([]byte(ident))
	str := strconv.Itoa(int(hash[0]))
	first := str[0:1]
	h, _ := strconv.Atoi(first)
	return h
}
