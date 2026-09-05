package main

import "fmt"

func main() {
	var x int
	x++
	fmt.Println(x)
	i := c()
	fmt.Println(i)
}

func c() (i int) {
	// i starts with its zero value: 0
	defer func() { i++ }() // The function literal is evaluated into a closure
	// and registered; its body is not executed yet.
	return 1 // Assigns 1 to i; then the deferred closure executes,
	// increments i to 2, and c returns 2.
}
