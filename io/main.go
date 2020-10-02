package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, error := os.Open("sample.txt")
	defer file.Close()
	if error != nil {
		log.Fatal(error)
	}

	reader := bufio.NewReader(file)
	total := 0
	for {
		byte, _, error := reader.ReadLine()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatal(error)
		}
		line := string(byte)
		c := strings.Count(line, "寺子屋")
		total += c
	}
	fmt.Println(total)
}
