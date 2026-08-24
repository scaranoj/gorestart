package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("myfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Go Broncos")
	fmt.Println(r)

	io.Copy(f, r)

}
