package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		if i/2 == 0 {
			fmt.Println("skipping even #")
			continue
		}
		if i%2 != 0 {
			fmt.Printf("The odd remainder is %v\n", i)
		}
	}
}
