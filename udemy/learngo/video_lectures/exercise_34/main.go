package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for i := 0; i < 5; i++ {
			fmt.Printf("inner loop %v\n", i)
		}
		fmt.Printf("outer loop %v\n", i)

	}
}
