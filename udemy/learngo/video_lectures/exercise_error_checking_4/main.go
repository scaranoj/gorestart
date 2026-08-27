package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	sb, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))
}
