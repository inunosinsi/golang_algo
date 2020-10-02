package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	url := "http://localhost:9999/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dump))
}
