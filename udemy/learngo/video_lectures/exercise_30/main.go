package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i <= 42; i++ {
		fmt.Printf("iteration %v\t", i)
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("x is %v\n", x)
		case 1:
			fmt.Printf("x is %v\n", x)
		case 2:
			fmt.Printf("x is %v\n", x)
		case 3:
			fmt.Printf("x is %v\n", x)
		case 4:
			fmt.Printf("x is %v\n", x)
		}
	}
}
