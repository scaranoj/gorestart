package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var x = rand.Intn(10)
	var y = rand.Intn(10)
	fmt.Printf("x is %v and y is %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or equal to 4 and less than 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the prev cases were met")
	}
}
