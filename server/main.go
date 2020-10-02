package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	// ブラウザでhttp://localhost:8080/と打ち込んだ時にhandler関数が実行される
	http.HandleFunc("/", handler)

	// ブラウザを開き、アドレスバーにhttp://localhost:8080と打ち込むと何らかを表示するようにする
	http.ListenAndServe(":9999", nil)
}

// http://localhost:8080/を開いた時に表示させたりする
func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	//端末にリクエストの内容が出力されます
	fmt.Println(string(dump))
}
