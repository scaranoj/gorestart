package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//preceding the if condition with a statement statement idiom to scope the variable to the if block

	//The number of lines in output is how many times Intn (psuedo)randomly rolled a 3
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("at iteration %v\t x is 3\n", i)
		}
	}
}
