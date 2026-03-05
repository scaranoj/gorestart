package main

import "fmt"

func main() {

	xs := make([]string, 0, 50)

	xs = append(xs, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Printf("The 50 states are %v:\n", xs)
	fmt.Printf("The length of the slice is %v:\n", len(xs))
	fmt.Printf("The capacity of the slice is %v\n", cap(xs))

	// if can't use a range, then prob needs to be a nested for loop
	//also, you could prob use the len() in here somewhere to get the index
	//but how to get extract the value at each index?
	//the slice operator could maybe do this (e.g. xs[i])
	// a range would look like this: for i, v := range xs

	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i], i)
	}
}

//my original loop (nested)
/*for i := 0; i < len(xs); i++ {
	fmt.Printf("When index is %v ", i)
	for v := 0; v <= 0; v++ {
		fmt.Printf("value is %v\n", xs[i])
	}

}*/

/* ● Create a slice to store the names of all of the states in the United States of America.
○ Use make and append to do this.
○ Goal: do not have the array that underlies the slice created more than once.
● Print out
○ the len
○ the cap
○ the values, along with their index position, without using the range clause.
● Here is a list of the 50 states:
` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
https://go.dev/play/p/-cwlxRpnJAw
*/
