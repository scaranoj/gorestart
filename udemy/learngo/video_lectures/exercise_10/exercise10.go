package main

import "fmt"

// var for zero value
var simpleString string = ""

// var when specifity is required
var c int = 32

func main() {
	// short declaration operator, can only be used within func body
	y := 5
	//multiple initializations
	a, b := 42, 32
	// blank identifier (i.e. for a value you don't need, like /dev/null garbage can)
	h, i, _ := 1, 2, 3
	fmt.Printf("The zero value of simpleString is: %v\n", simpleString)
	fmt.Printf("The value of y via short declaration is %v\n", y)
	fmt.Printf("The values of a and b using multiple initializations is: %v and %v\n", a, b)
	fmt.Printf("The value of c via the var identifier is %v\n", c)
	fmt.Printf("The value of blank identifier _ is %v %v\n", h, i)

}
