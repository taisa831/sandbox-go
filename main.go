package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("sample.txt")
	b := make([]byte, 1)
	for {
		n, err := file.Read(b)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}
}
