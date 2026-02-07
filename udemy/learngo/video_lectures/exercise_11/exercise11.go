package main

import "fmt"

func main() {

	x, y, z := "Bears", 5, 23.22

	fmt.Printf("%v \t %T\n", x, x)

	fmt.Printf("%v \t %T\n", y, y)

	fmt.Printf("%v \t %T\n", z, z)

	/*
		x := "Bears"
		fmt.Printf("%v \t %T\n", x, x)
		y := 5
		fmt.Printf("%v \t %T\n", y, y)
		z := 23.22
		fmt.Printf("%v \t %T\n", z, z)
	*/
}

/*
write a program that uses the following:
● for a VARIABLE storing a VALUE of TYPE
	○ string
	○ int
	○ float64
● use print verbs to show
	○ value
	○ type
*/
