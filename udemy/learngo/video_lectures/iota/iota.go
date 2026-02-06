package main

import "fmt"

// before using iota, need to define a type here based on int that will represent
// all valid values
type c int

const (
	c0 c = iota // c0 == 0
	c1          // c1 == 1 - no need to specify type or value for the rest of these constants
	c2          // c2 == 2
)

const (
	c3 c = iota // c0 == 0 - new block resets iota back to value (index) of 0
	c4          // again, the rest of this block's constants neither needs a type or value assigned to them
	c5
	c6
)

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Println(c3, c4, c5, c6)
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)
}
