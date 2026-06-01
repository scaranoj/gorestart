package main

// Create a value and assign it to a variable.
// Print the address of that value

import "fmt"

var a int

func main() {

	fmt.Println(&a, "", a)
}
