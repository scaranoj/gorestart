package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Println("Error happened", err)
	}
	defer f.Close()
}
