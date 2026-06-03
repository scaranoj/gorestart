// * In the provided code, there are variables that store VALUES of TYPE *int and TYPE *string. The values stored are memory addresses. Using the variables provided, do the following:
// * print the VALUE stored in each variable
//   * these will be memory addresses
// * print the TYPE of each variable
// * print the data stored at memory locations
//   * dereference the stored memory address *

package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {

	fmt.Printf("The value of a is: %v, the type is %T, and the value at the memory location is: %v\n", a, a, *a)
	fmt.Printf("The value of b is: %v, the type is %T, and the value at the memory location is: %v\n", b, b, *b)
	fmt.Printf("The value of c is: %v, the type is %T, and the value at the memory location is: %v\n", c, c, *c)
	fmt.Printf("The value of d is: %v, the type is %T, and the value at the memory location is: %v\n", d, d, *d)

}
