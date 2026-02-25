package main

import (
	"fmt"
)

func main() {

	i := 10

	for {
		if i < 0 {
			break
		}
		fmt.Printf("%v\n", i)
		i--
	}
}
